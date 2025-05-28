# 📝 Go CLI Todo List

A simple command-line todo list manager written in Go. Tasks are stored in a local text file (`todo.txt`), and you can add, list, or mark them as done.

---


## ⚙️ How It Works

- The application is split into two files:  
  - `main.go`: Handles command-line input and routes to the correct function using a switch-case.
  - `todo.go`: Contains the actual functions (`AddTodo`, `ListTodos`, `MarkDone`) that interact with the `todo.txt` file.

- Tasks are stored in a simple text file (`todo.txt`), one task per line:
  - Incomplete tasks are saved as: `[ ] Task description`
  - Completed tasks are saved as: `[x] Task description`

---

## 🚀 Usage

### ▶️ Run the app

Make sure you're in the project folder and have Go installed.

```bash
go run main.go todo.go [command] [arguments...]