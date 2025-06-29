# 🏗️ Go Project: Three-Layer Architecture (Clean Code Example)

This project demonstrates a **Three-Layer Architecture** in Golang with unit testing and mock integration, using:

- `net/http` for handlers
- A custom `service` layer for business logic
- A `store` layer for data access
- Fully unit-tested using Go’s `testing` package and mock functions

---

## 🧱 Project Structure

microservice/
│
├── Models/
│   ├── user/
│   │   ├── user.go
│   │   └── user_test.go              # ✅ model-related tests (if any)
│   └── task/
│       ├── task.go
│       └── task_test.go
│
├── Store/
│   ├── user/
│   │   ├── store.go
│   │   └── store_test.go             # ✅ user store tests
│   └── task/
│       ├── store.go
│       └── store_test.go             # ✅ task store tests
│
├── Service/
│   ├── user/
│   │   ├── service.go
│   │   └── service_test.go           # ✅ service logic test
│   └── task/
│       ├── service.go
│       └── service_test.go
│
├── Handler/
│   ├── user/
│   │   ├── handler.go
│   │   └── handler_test.go           # ✅ HTTP handler test
│   └── task/
│       ├── handler.go
│       └── handler_test.go
│
├── Database/
│   ├── connection.go
│   └── connection_test.go           # ✅ DB connection test
│
└── main.go
