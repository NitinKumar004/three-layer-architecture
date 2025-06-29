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

//no issue golangci-lint run checked
zopdev@ZopSmarts-MacBook-Pro microservice % golangci-lint run                                                              
0 issues.

// test  coverge
go test ./... --coverprofile=c.out                                             
microservice            coverage: 0.0% of statements
?       microservice/Models/task        [no test files]
?       microservice/Models/user        [no test files]
ok      microservice/database   2.033s  coverage: 0.0% of statements [no tests to run]
ok      microservice/handler/task       1.496s  coverage: 73.2% of statements
ok      microservice/handler/user       1.237s  coverage: 71.4% of statements
ok      microservice/service/task       0.936s  coverage: 84.6% of statements
ok      microservice/service/user       1.752s  coverage: 82.4% of statements
ok      microservice/store/task 0.512s  coverage: 88.1% of statements
ok      microservice/store/user 2.360s  coverage: 82.5% of statements

