# Go Contact Book

This is a project I'm using to learn the Go language. It started off as a CLI tool, and the plan is to turn it into a backend HTTP server with database connections and online deployment. The goal here is to learn Golang at least to a basic level.

## Features Implemented (So Far)

- Implement fully working contact writing, reading and lookup via CLI arguments (deprecated)
- Implement JSON file storage for contact persistence (deprecated)
- Implement all CRUD operations related to contacts
- Implement RESTFUL endpoints for executing CRUD operations
- Implement **Repository Pattern** to ensure **Separation of Concerns** between application layers.
- Implement **unit tests** for the Service layer functions.
- Implement **SQLite** database storage for robust persistence.

---

## 🚀 Setup and Run

This project uses Go for the server logic and **SQLite** for persistence.

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.21 or later)

### 1. Clone the Repository

```bash
git clone <your_repo_url>
cd go-contact-book
```

### 2. Install Dependencies

You need to install the database driver required by the project.

```bash
go get [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
go mod tidy
```

### 3. Run the Server

The application will automatically create the contacts.db file in the project directory on the first run.

```bash
go run main.go
```

![Go Logo](https://img.shields.io/badge/golang-00ADD8?&style=plastic&logo=go&logoColor=white)
![SQLite Logo](https://img.shields.io/badge/SQLite-4169E1?logo=sqlite&logoColor=fff&style=plastic)
