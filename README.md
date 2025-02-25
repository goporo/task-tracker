# Task Tracker CLI

Task Tracker CLI is a simple command-line application to help you manage your tasks efficiently. It allows you to add, update, delete, mark, and list tasks.

## Features

- ✅ Add new tasks
- ✏️ Update task descriptions
- ❌ Delete tasks
- 📌 Mark tasks with statuses: `todo`, `in-progress`, or `done`
- 📋 List all tasks or filter by status

## Installation

1. Clone the repository:
```sh
git clone https://github.com/yourusername/task-tracker.git
cd task-tracker
```

2. Build the application:
```sh
go build -o task-tracker ./cmd/task-tracker
```

## Usage

### Add a new task
```sh
./task-tracker add "New Task Description"
```

### Update a task description
```sh
./task-tracker update 1 "Updated Task Description"
```

### Delete a task
```sh
./task-tracker delete 1
```

### Mark a task as done
```sh
./task-tracker mark 1 done
```

### List all tasks
```sh
./task-tracker list
```

### List tasks by status
```sh
./task-tracker list todo
```

### List tasks by optional status
If provided, the `[status]` parameter filters tasks by the given status (`todo`, `in-progress`, or `done`).

```sh
./task-tracker list [status]
```

## Project Structure

- `cmd`
  - `root.go`: Defines the CLI commands and their handlers.
- `internal/tasks`
  - `task.go`: Contains the logic for managing tasks (loading, saving, adding, updating, deleting, marking, and listing tasks).
- `main.go`: Entry point of the application.
- `tasks.json`: JSON file to store tasks.

## License

This project is licensed under the MIT License.
