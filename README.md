# To-Do List REST API (Go)

This project is a simple, educational RESTful API for managing to-do tasks, built with Go. It demonstrates how to build a backend service using Go's standard library, SQLite for persistence and includes a web-based frontend for easy interaction.

## Project Overview

- **Backend:** Go (net/http, encoding/json, database/sql, SQLite)
- **Frontend:** HTML/JavaScript (served at `/`)
- **API Spec:** OpenAPI (`openapi.yaml`)
- **Demo:** Add, update, delete and list tasks via API or web UI

## Why This Project?

This project is ideal for:

- Showcasing Go skills in web/API development
- Learning RESTful API design and CRUD operations
- Demonstrating SQLite integration in Go
- Providing a full-stack demo for portfolio or interviews

## Features

- CRUD operations for tasks (Create, Read, Update, Delete)
- Persistent storage with SQLite
- Simple web frontend for non-technical users
- OpenAPI spec for easy API exploration

## Tech Stack

- Go 1.21+
- SQLite (via `github.com/mattn/go-sqlite3`)
- HTML/JavaScript (for demo UI)

## Setup & Usage

### Prerequisites

- Go installed (`go version`)
- SQLite (no manual setup needed; file created automatically)

### Running the API

1. Clone the repository:

```bash
  git clone https://github.com/yourusername/todo-api.git
  cd todo-api
```

2. Install dependencies:

```bash
  go mod tidy
```

3. Start the server:

```bash
  go run main.go
```

4. Open [http://localhost:8080/](http://localhost:8080/) in your browser for the demo UI.

### API Endpoints & Examples

### Create a Task

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"Test Task","completed":false}' http://localhost:8080/tasks
```

**Response:**

```json
{
  "id": 1,
  "title": "Test Task",
  "completed": false
}
```

### List All Tasks

**Request:**

```bash
curl http://localhost:8080/tasks
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "Test Task",
    "completed": false
  }
]
```

### Get a Single Task

**Request:**

```bash
curl http://localhost:8080/tasks/1
```

**Response:**

```json
{
  "id": 1,
  "title": "Test Task",
  "completed": false
}
```

### Update a Task

**Request:**

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Task","completed":true}' http://localhost:8080/tasks/1
```

**Response:**

```json
{
  "id": 1,
  "title": "Updated Task",
  "completed": true
}
```

### Delete a Task

**Request:**

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

**Response:**
HTTP 204 No Content

## API Documentation

See [`openapi.yaml`](openapi.yaml) for a full OpenAPI spec. You can use [Swagger UI](https://swagger.io/tools/swagger-ui/) or [Redoc](https://redocly.com/) to visualize and test the API.

## Frontend Demo

Open [`frontend.html`](frontend.html) in your browser, or visit [http://localhost:8080/](http://localhost:8080/) while the server is running.

## Contributing

Pull requests and suggestions are welcome! Feel free to fork, improve, or open issues.

## License

MIT License. See [`LICENSE`](LICENSE) for details.

## Contact

Created by oboboss11. For questions, reach out via GitHub Issues or your preferred contact method.

## Example Task JSON

```json
{
  "title": "Learn Go",
  "completed": false
}
```
