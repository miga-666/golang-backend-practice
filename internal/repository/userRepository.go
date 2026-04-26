package repository

import (
	"ginBackend/internal/model"

	"gorm.io/gorm"
)

// Defines the data access operations for the User domain.
type UserRepository interface {
	// Looks up a user by their email address.
	FindByEmail(email string) (*model.User, error)

	// Inserts a new user record
	Create(user *model.User) error

	// Update the hashed password by user's email
	UpdatePassword(email string, hashedPassword string) error
}

// GORM-backed implementation of UserRepository.
type userRepository struct {
	db *gorm.DB
}

// Constructs a UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Queries the database by the user's email
// Returns the User struct if found
func (repo *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := repo.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create inserts the given User struct as a new row
func (repo *userRepository) Create(user *model.User) error {
	return repo.db.Create(user).Error
}

// Updates the 'password' and 'updated' column by the user's email
func (repo *userRepository) UpdatePassword(email string, hashedPassword string) error {
	return repo.db.Model(&model.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"password": hashedPassword,
			"updated":  gorm.Expr("NOW()"),
		}).Error
}
