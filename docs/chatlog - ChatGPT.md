面試前有個 golang 專案測驗
Golang RESTful API 實作題目
題目說明
請使用 Golang 建立一個簡易的 RESTful API 後端服務,並串接關聯式資料庫,實現以下功能:
0. API 使用統一格式回覆
所有 API 皆須使用以下統一格式回覆:
{
"success": true,
"message": "成功",
"data": {
// 回傳資料
},
"error": null,
"code": 200,
"timestamp": "2025-09-12T07:58:43Z"
}
{
"success": false,
"message": "失敗原因",
"data": null,
"error": {
// 詳細原因

},
"code": 400,
"timestamp": "2025-09-12T07:58:43Z"
}

1. 使用者註冊
- 提供 API 讓新使用者可以註冊帳號。
- 註冊時需儲存使用者資料(帳號、密碼)至資料庫,密碼使用 bcrypt 加密處理。
- 若 email 已存在,需回傳對應的錯誤訊息。
- 不需要 JWT Token

2. 使用者登入
- 提供 API 讓使用者可以登入。
- 登入時需驗證帳號與密碼,並回傳適當的驗證結果。
- 成功時回傳 JWT Token(payload 須包含 email 與 updated 欄位),失敗時回傳錯誤原
因。
- 帳號或密碼錯誤時,錯誤訊息不應揭露究竟是帳號還是密碼有誤。
- 不需要 JWT Token

3. 修改密碼
- 提供 API 讓使用者可以修改自己的密碼。
- 參數需包含舊密碼與新密碼。
- 修改後需將新密碼(bcrypt 加密後)更新至資料庫。
- 修改成功後,原有的 JWT Token 仍然有效,不需實作 Token 黑名單機制。
- 需要 JWT Token(請在 Request Header 帶上 Authorization: Bearer <token>)

4. User 資料表
column data type PK 說明
email varchar(50) ✅ 使用者帳號,唯一識

別鍵

password varchar(255) bcrypt 加密後的密碼

雜湊值
created datetime 帳號建立時間
updated datetime 最後更新時間

5. 如有使用 AI 工具輔助開發
請提供以下兩份文件:
1. SPEC 文件:對整個專案的規劃說明,需為結構化文件(如 Markdown),需詳細到可以讓
AI 工具直接讀取並理解你的開發意圖,進而協助你完成實作。
2. AI 對話紀錄:與 AI 工具的實際互動紀錄(截圖或文字均可),需能呈現你如何運用 AI 協助
開發的過程。
以上兩份文件請放置於 repo 中(建議統一放在 docs/ 資料夾)。

6. 加分項目
1. 串接中央氣象署 API,獲取資料後,只取新北市的資料儲存進 DB。
- API:/api/v1/rest/datastore/F-C0032-001 中央氣象署開放資料平臺之
資料擷取API
- 請至 中央氣象署開放資料平台 申請免費 API Key,並將其設定於環境變數中,勿
寫入程式碼。
- 排程需在服務啟動時立即執行一次,並每 24 小時重新獲取並更新一次。
2. Weather 資料表

column data type PK 說明
id BIGINT / SERIAL ✅ 主鍵,自動流水號
city VARCHAR(50) 城市名稱(例如:新北

市)

min_t DECIMAL(4,1) 最低氣溫 (°C)
max_t DECIMAL(4,1) 最高氣溫 (°C)
period ENUM('AM','PM') 上午 (06:00–12:00) /
下午 (12:00–18:00)

created datetime 建立時間
updated datetime 更新時間
3. 新增兩個天氣查詢 API,回傳新北市今日上午(06:00–12:00)與下午(12:00–18:00)的氣
溫資料,皆需使用統一格式回覆:
- 一個 API 不需要 JWT Token(任何人可存取)
- 一個 API 需要 JWT Token(Authorization: Bearer <token>)

參考資料:https://pjchender.dev/react-bootcamp/docs/book/ch5/5-1/

必要條件
- 請使用 Gin 框架實作 API。
- 資料庫請使用 MySQL。
- 所有 API 請遵循 RESTful 設計原則,路由設計由你自行決定。
- 提供資料表的 CREATE 語法,或者 golang-migrate 指令。
- 請撰寫 README,提供在本機端啟動服務的完整指令說明。
- 請提供 Dockerfile 與 docker-compose,需同時包含 App container 與 MySQL
container,確保可用單一指令在本機啟動完整服務。
- 請提供 API 文件(以下擇一),且必須可以在本機發送 request 給本機服務:
- Swagger UI(工具自動產生即可)
- Postman(可匯入的 collection.json)

需要提供以下東西
- GitHub 連結(請將 repo 設為 Public,以供審閱)
Mysql 和 go 我都沒用過 也沒從頭建過傳案後端serverc和 sql server 現在需要先搞清楚關於專案的一些地方
1. 介紹 GO
2. 完完整整介紹 GIN, 並推薦我 YT 教學影片
3. 開發工具用哪個 (VS/VS CODE)
4. AI 用 chatGPT / Copilot
5. 介紹 docker container (backend + DB 都 run 在 container 裡嗎), docker file 怎麼寫, 有哪寫 run container 的指令
6. 如何建立 DATABASE 和 Table, golang-migrate 指令是什麼怎麼用
7. Gin 的 router 和 restful api 怎麼寫
8. 簡單介紹 bcrypt 驗算法, go 是否有套件
9. 什麼是 SPEC DOC? 給我一個實際慶況的範例, 用文字呈現即可
10. 解釋題目規定的 API 統一格式, 以及 peyload 是什麼
11. Gin 和 MySQL 連接的方式
12. 需要 Dapper 嗎? 順便介紹一下 Dapper 是什麼
13. MySQL 和 MS SQL 的差異
14. JWT token 的原理與如何實現
15. 需要把 controller 和 DAO 分開嗎
16. 好的 SPEC 需要哪些資訊
17. 如果要將 backend and db 部屬到 container 裡, 開發時就要 CICD 了嗎? db 和 table 怎麼建? 還是說可以先在本地開發好, 再包進 docker?
幫我畫完整 Go + Gin 專案架構（資料夾 + 每個檔案內容）
用 markdown 基於題目(含加分題)幫我寫一個 spec document 範例

我還有些額外要求

變數名稱用駝峰式命名 不能用單個字母單變數 request 就用 req repoonse 就用 res
要寫註解
要分層 controllor service repo 物件 等等等
我決定先把 weather 部分移除 並修改了一下順序 以及 API 的部分 幫我完善他
# Go + Gin Backend Service SPEC Document

[TOC]

## 1. Project Overview

This project is a RESTful backend service built with Go (Gin framework) and MySQL.

It provides:
- User authentication (register / login)
- Password management
- JWT-based authorization
- Standardized API response format
- Dockerized deployment environment


## 2. Tech Stack
- Language: Go
- Web Framework: Gin
- Database: MySQL
- ORM (optional): GORM
- Authentication: JWT
- Password Hashing: bcrypt
- Container: Docker + Docker Compose
- Migration Tool: golang-migrate

## 3. Architecture Design

### 3.1 Layered Architecture
- Controller → Service → Repository → Database


### Responsibilities:

| Layer | Responsibility |
|------|--------|
| Controller | Handle HTTP request/response |
| Service | Business logic |
| Repository | Database operations |
| Model | Data structure |
| Middleware | JWT validation |
| Utils | Common helpers |


## 4. Naming Convention Rules

### MUST FOLLOW:

- Use **camelCase**
- NO single-letter variable names
- request object must use req
- response object must use res

### Example:


go
// GOOD
var loginRequest LoginRequest
var loginResponse LoginResponse

// BAD
var r LoginRequest
var x LoginResponse



## 6. Database Schema

### User Table


sql
CREATE TABLE users (
    email VARCHAR(50) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    created DATETIME,
    updated DATETIME
);




## 5. API Standard Response Format

### Success Response


json
{
  "success": true,
  "message": "成功",
  "data": {},
  "error": null,
  "code": 200,
  "timestamp": "2026-01-01T00:00:00Z"
}


### Error Response


json
{
  "success": false,
  "message": "錯誤原因",
  "data": null,
  "error": {
    "detail": "internal error"
  },
  "code": 400,
  "timestamp": "2026-01-01T00:00:00Z"
}


## 14. Error Handling Rules

* Never expose internal errors to client
* Always return standardized error format
* Authentication failure must not reveal whether email or password is wrong





## 7. API Design


## 7.1 Register API

### Endpoint


POST /api/v1/auth/register


### Request


json
{
  "email": "test@gmail.com",
  "password": "123456"
}


### 邏輯?
- 註冊時需儲存使用者資料（帳號、密碼）至資料庫

### Flow
- 密碼使用 bcrypt 加密處理
- 不需要 JWT Token
- success case: 新增帳號成功: email
- fail case:
    - 未輸入 email
    - 未輸入 password
    - email 已存在 



## 7.2 Login API

### Endpoint


POST /api/v1/auth/login


### Request

json
{
  "email": "test@gmail.com",
  "password": "123456"
}

### Flow
- 讓使用者可以登入
1. Validate user input
2. Find user by email
3. Compare bcrypt password
4. Generate JWT token
- success case: 登入成功: email, 回傳 JWT Token
- fail case:
    - 未輸入 email
    - 未輸入 password
    - 帳號或密碼錯誤

## 7.3 修改密碼 API
### Request

json
{
  "email": "test@gmail.com",
  "oldPassword": "123456",
  "newPassword": "654321"
}




- 提供 API 讓使用者可以修改自己的密碼。
- 參數需包含舊密碼與新密碼。
- 修改後需將新密碼（bcrypt 加密後）更新至資料庫。
- 修改成功後，原有的 JWT Token 仍然有效，不需實作 Token 黑名單機制。
- 需要 JWT Token（請在 Request Header 帶上 Authorization: Bearer <token>）
- success case: 修改密碼成功
- fail case:
    - 未登入或 token 失效
    - 帳號或舊密碼錯誤
    - 未輸入 email
    - 未輸入新 password
    - 未輸入舊 password

## 13. Environment Variables


DB_HOST=ginUser
DB_PASSWORD=p@ssword
DB_NAME=ginBackend




## 12. Docker Setup
- docker-compose.yml
    - backend
    - db


    
    
## 18. Deployment Requirement

Project must be runnable via:


bash
docker-compose up --build


and should include:

* backend container
* mysql container
* automatic database connection

JWT Token（請在 Request Header 帶上 Authorization: Bearer <token>） 這是什麼意思

JWT token 可以在 gin 如何實現? 用前端和後端互相傳遞的實際例子給我

解釋以下幾個 layer 的用途 排呈要放哪 為何 JWT 是 middleware Utils 	是什麼 Repository 	要寫通用的還是不通用的
Controller 	Handle HTTP request/response
Service 	Business logic
Repository 	Database operations
Model 	Data structure
Middleware 	JWT validation
Utils 	Common helpers
現在我有一份 SPEC Document 
已經建好一個空資料夾 db 也建了
如何用 copilot agent 幫我 coding

註: 且我不想直接包進 docker 想先測完 API 再包
我再 doc 裡再加一規則
程式 function 的順序 閱讀的起始點是最下面 由下往上
 

A()

B1()

B2()

B() {
B1
B2
}
main() {
A()
B()

}
package repository

import (
	"ginBackend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	UpdatePassword(email string, hashedPassword string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := repo.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *userRepository) Create(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepository) UpdatePassword(email string, hashedPassword string) error {
	return repo.db.Model(&model.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"password": hashedPassword,
			"updated":  gorm.Expr("NOW()"),
		}).Error
}


解釋

func (svc *userService) Register(req *model.RegisterRequest) (*model.RegisterResponse, error) {
	_, err := svc.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("database error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	now := time.Now()
	user := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Created:  now,
		Updated:  now,
	}

	if err := svc.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return &model.RegisterResponse{Email: req.Email}, nil
}

這一段會包 APIResponse 嗎
return
在這專案需要 middleware 嗎
我已在本機 run backend 和 db 確認程式和 API 正確 如何改成可以運行在 docker  的版本?


