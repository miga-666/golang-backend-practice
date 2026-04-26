# SPEC Document

## 1. Project Overview

This project is a RESTful backend service built with Go (Gin framework) and MySQL.

It provides:

- User registration
- User login
- Password update
- JWT authentication
- Standard API response format
- Dockerized deployment (backend + database)

## 2. Tech Stack

- Language: Go
- Web Framework: Gin
- Database: MySQL
- Authentication: JWT
- Password Hashing: bcrypt
- Container: Docker + Docker Compose
- Migration Tool: golang-migrate (optional)

## 3. Architecture Design

### 3.1 Layered Architecture

```

Controller → Service → Repository → Database

````

### 3.2 Layer Responsibilities

| Layer | Responsibility |
|------|--------|
| Controller | Handle HTTP request/response |
| Service | Business logic |
| Repository | Database operations |
| Model | Data structure |
| Middleware | JWT validation |
| Utils | Common helpers |

## 4. Naming Convention Rules

### MUST FOLLOW

- Use camelCase
- No single-letter variables
- Request object must use `req`
- Response object must use `res`

### Example

```go
var loginReq LoginRequest
var loginRes LoginResponse
````

## 5. API Standard Response Format

### Success Response

```json
{
  "success": true,
  "message": "成功",
  "data": {},
  "error": null,
  "code": 200,
  "timestamp": "2026-01-01T00:00:00Z"
}
```

### Error Response

```json
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
```

## 6. Database Schema

### 6.1 User Table

```sql
CREATE TABLE users (
    email VARCHAR(50) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    created DATETIME,
    updated DATETIME
);
```

## 7. Authentication Design (JWT)

### 7.1 JWT Payload

```json
{
  "email": "user@gmail.com",
  "updated": "2026-01-01T00:00:00Z"
}
```

### 7.2 JWT Usage Rules

* Login API returns JWT token
* Protected APIs require:

```
Authorization: Bearer <token>
```

* Token validation is handled by middleware

## 8. API Design

# 8.1 Register API

### Endpoint

```
POST /user/register
```

### Request

```json
{
  "email": "test@gmail.com",
  "password": "123456"
}
```

### Flow

1. Validate request
2. Check email exists
3. Hash password using bcrypt
4. Insert into DB

### Success Response

```json
{
  "success": true,
  "message": "註冊成功",
  "data": {
    "email": "test@gmail.com"
  },
  "error": null,
  "code": 200
}
```

### Failure Cases

* email missing
* password missing
* email already exists

# 8.2 Login API

### Endpoint

```
POST /user/login
```

### Request

```json
{
  "email": "test@gmail.com",
  "password": "123456"
}
```

### Flow

1. Validate request
2. Find user by email
3. Compare bcrypt password
4. Generate JWT token

### Success Response

```json
{
  "success": true,
  "message": "登入成功",
  "data": {
    "email": "test@gmail.com",
    "token": "jwt_token_here"
  },
  "error": null,
  "code": 200
}
```

### Failure Cases

* invalid email or password (do NOT reveal which one is wrong)
* missing input fields

# 8.3 Change Password API

### Endpoint

```
PUT /user/password
```

### Auth Required

```
Authorization: Bearer <token>
```

### Request

```json
{
  "email": "test@gmail.com",
  "oldPassword": "123456",
  "newPassword": "654321"
}
```

### Flow

1. Validate JWT token
2. Verify email from token OR request
3. Check old password
4. Hash new password
5. Update DB

### Success Response

```json
{
  "success": true,
  "message": "密碼已更新",
  "data": null,
  "error": null,
  "code": 200
}
```

### Failure Cases

* unauthorized (invalid token)
* email not found
* old password incorrect
* missing fields

## 9. Security Rules

* Password must be hashed using bcrypt
* JWT secret must not be hardcoded
* Do not expose internal errors to client
* Login error must not reveal whether email or password is incorrect

## 10. Environment Variables

```env
DB_HOST=db
DB_USER=root
DB_PASSWORD=root
DB_NAME=gin_backend
JWT_SECRET=your_secret_key
```

## 11. Docker Setup

### Services

* backend (Go + Gin)
* db (MySQL)

## 12. System Flow Overview

### Register Flow

```
Client → Controller → Service → Repository → DB
```

### Login Flow

```
Client → Controller → Service → Repository → DB → JWT → Response
```

### Password Update Flow

```
Client → JWT Middleware → Controller → Service → Repository → DB
```

## 13. Error Handling Rules

* Always return unified response format
* Never expose stack trace to client
* Authentication failure must be generic

## 14. Deployment Requirement

System must run via:

```bash
docker-compose up --build
```

Must include:

* backend container
* mysql container
* automatic DB connection
