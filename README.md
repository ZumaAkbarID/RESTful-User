# 🧑‍💻 Golang RESTful API - CRUD User

A simple RESTful API project for CRUD users, built with [Golang](https://go.dev/), [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/) (MySQL), and [go-playground/validator](https://github.com/go-playground/validator). Supports CORS and Docker for easy deployment.

---

## 🚀 Features

- CRUD operations for users
- Input validation using `go-playground/validator`
- CORS middleware support
- MySQL database integration using GORM
- Docker support with `.env` configuration
- API documentation using [Apidog](https://apidog.com/)
- 🔥 Live reload during development with [Air](https://github.com/cosmtrek/air)

---

## 📁 Project Structure

```yaml
.
├── main.go
├── go.mod
├── go.sum
├── Dockerfile
├── .air.toml
├── .env.example
├── .gitignore
├── .dockerignore
├── README.md
├── /config
├── /models
├── /controllers
├── /routes
├── /database
├── /helpers
├── /structs
└── /middlewares
```

---

## 🔧 Getting Started

### 1. Clone Repository

```bash
git clone https://github.com/ZumaAkbarID/RESTful-User.git
cd RESTful-User
```

### 2. Create Environment File

```bash
cp .env.example .env
```

#### Edit .env with your configuration

---

## 🐳 Running with Docker

### 1. Build Docker Image

```bash
docker build -t golang-user-api .
```

### 2. Run Docker

```bash
docker run --env-file .env -p 3000:3000 golang-user-api
```

## 📦 Running Locally (Without Docker)

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Create Environment File

```bash
cp .env.example .env
```

#### Edit .env with your configuration

### 3. Run App

```bash
air
```

---

## 📘 API Documentation

You can find and test the API documentation via Apidog:

[🔗 View API Documentation on Apidog](https://zm-go-restful-user.apidog.io)

| Method | Endpoint         | Description       |
| ------ | ---------------- | ----------------- |
| GET    | `/`              | Hello World       |
| GET    | `/api/users`     | Get all users     |
| GET    | `/api/users/:id` | Get user by ID    |
| POST   | `/api/users`     | Create new user   |
| PUT    | `/api/users/:id` | Update user by ID |
| DELETE | `/api/users/:id` | Delete user by ID |

## 📄 License

MIT © 2025 — [ZumaAkbarID](https://github.com/ZumaAkbarID)
