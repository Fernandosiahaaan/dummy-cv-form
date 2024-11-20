# dummy-cv-form
this repo describe to create dummy form cv application

# 📚 dummy-cv-form

this repo describe to create dummy form cv application

## Tech Stack 

This project utilizes several modern technologies and follows a clean architecture approach to ensure scalability, maintainability, and ease of development.

| Technology/Tool        | Purpose                                               | Status |
| ---------------------- | ----------------------------------------------------- | ------ |
| **Golang**             | Core language for the back-end services               | ✅     |
| **Gorilla Mux**        | HTTP router for handling requests                     | ✅     |
| **Monolith**           | Monolithic repository                                 | ✅     |
| **Clean Architecture** | Layered design for maintainability                    | ✅     |
| **Redis**              | Caching layer for improving performance               | ✅     |
| **PostgreSQL**         | SQL database for task and user services               | ✅     |
| **Docker**             | Containerization for setup and environment management | ✅     |

## 🖊 Pre Setup

- Install docker desktop 
- Install VSCode
- clone this project

## 🖊 Start/ Run project

Runing docker compose

```
cd /dummy-cv-form
docker-compose up -d
```

After success, run migration

```
cd /migration
go run .           # set up/down in migration.go line 31 based your need.
```

After success, run app.

```
cd /cmd
go run .
```
