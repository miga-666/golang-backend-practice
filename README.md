# README

## Run with Docker
```bash
docker-compose up --build
```
- Backend: `http://localhost:8080`


## API Testing with Postman
- Import:
  - `docs/postman - collection.json`
  - `docs/postman - environment.json`
- Then:
  - Select environment
  - Call protected APIs
  - Run `使用者登入` → save token manually to environment variable `token` 