# Go JWT Authentication API 🔑

A simple **Go** API using **Gin**, **PostgreSQL**, and **JWT** for user authentication.  
This project provides user registration, login, and protected routes with token-based authentication.

---

## 🚀 Features
- User registration with **hashed passwords** (bcrypt)
- User login with **JWT token generation**
- Middleware to protect routes
- PostgreSQL integration
- Environment variable-based configuration

---

## 📂 Project Structure
golang-jwt-authentication/
├── main.go # Entry point
├── middleware/ # JWT middleware
├── models/ # Database models
├── handlers/ # API handlers (register, login, etc.)
├── database/ # DB connection setup
├── .env # Environment variables (not committed)
├── go.mod # Go module file
└── README.md # Project documentation


---

## ⚙️ Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/your-username/golang-jwt-authentication.git
cd golang-jwt-authentication


### Install Dependencies
go mod tidy



### 3. Configure environment variables

Create a `.env` file in the project root with the following content:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432 # this is the predefined port number
DB_USER=postgres
DB_PASSWORD= # password 
DB_NAME=go_auth #DB name
SSL_MODE=disable
JWT_SECRET=your-secret-key

