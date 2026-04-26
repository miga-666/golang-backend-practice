package service

import (
	"errors"
	"time"

	"ginBackend/internal/model"
	"ginBackend/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	// Creates a new user account
	Register(req *model.RegisterRequest) (*model.RegisterResponse, error)

	// Login and returns a JWT on success
	Login(req *model.LoginRequest) (*model.LoginResponse, error)

	// Canges the user's password
	ChangePassword(req *model.ChangePasswordRequest) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Creates a new user account
func (svc *userService) Register(req *model.RegisterRequest) (*model.RegisterResponse, error) {
	// Check whether the email is already in use.
	_, err := svc.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("電子郵件已被使用")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("資料庫錯誤")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密碼加密失敗")
	}

	now := time.Now()
	user := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Created:  now,
		Updated:  now,
	}

	if err := svc.userRepo.Create(user); err != nil {
		return nil, errors.New("建立使用者失敗")
	}

	return &model.RegisterResponse{Email: req.Email}, nil
}

// Login and returns a JWT on success.
func (svc *userService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := svc.userRepo.FindByEmail(req.Email)
	if err != nil {
		// Do not reveal whether the email exists or not.
		return nil, errors.New("電子郵件或密碼錯誤")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("電子郵件或密碼錯誤")
	}

	token, err := generateJWT(user.Email, user.Updated)
	if err != nil {
		return nil, errors.New("產生驗證令牌失敗")
	}

	return &model.LoginResponse{Email: user.Email, Token: token}, nil
}

// Canges the user's password
func (svc *userService) ChangePassword(req *model.ChangePasswordRequest) error {
	user, err := svc.userRepo.FindByEmail(req.Email)
	if err != nil {
		return errors.New("找不到此電子郵件")
	}

	// Confirm the caller knows the current password before allowing a change.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("舊密碼不正確")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密碼加密失敗")
	}

	return svc.userRepo.UpdatePassword(req.Email, string(hashedPassword))
}
