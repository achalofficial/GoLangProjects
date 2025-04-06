# 🔗 Go URL Shortener

A simple, fast, and extensible URL shortening service built with Go, Gin, and PostgreSQL.

## 🛠 Tech Stack
- GoLang
- Gin Web Framework
- PostgreSQL

## 🚀 Features
- Create short URLs
- Redirect to original URLs
- Optional expiry support
- Custom aliases

## 📦 Local Setup

```bash
git clone https://github.com/yourusername/1-go-url-shortener.git
cd 1-go-url-shortener
go run cmd/main.go
```

Make sure PostgreSQL is running and you’ve set .env with:

ini
Copy
Edit
DB_URL=postgres://username:password@localhost:5432/yourdb
BASE_URL=http://localhost:8080


---

### ✅ Git Setup

From the root of the project folder:

```bash
git init
git add .
git commit -m "Initial working version of URL shortener"
gh repo create 1-go-url-shortener --public --source=. --remote=origin --push
