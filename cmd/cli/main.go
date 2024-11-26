package main

import (
	"fmt"
	"os"
	"todo-cli/infrastructure/database"
	"todo-cli/internal/repository"
	"todo-cli/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.NewSQLiteConnection("tasks.db")
	defer db.Close()

	database.Migrate(db)

	taskRepo := repository.NewTaskRepository(db)

	addTask := usecase.NewAddTask(taskRepo)
	getAllTask := usecase.NewGetAllTask(taskRepo)

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [<args>]")
		fmt.Println("Available commands:")
		fmt.Println("  add <title> <description> - Add a new task")
		fmt.Println("  list - List all tasks")
		fmt.Println("  complete <id> - Mark a task as completed")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo add <title> <description>")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		if err := addTask.Execute(title, description); err != nil {
			fmt.Println("Failed to add task:", err)

		}
		fmt.Println("Task added successfully")

	case "list":
		tasks, err := getAllTask.Execute()
		if err != nil {
			fmt.Println("Failed to list tasks:", err)
			return
		}
		fmt.Println("Tasks:")
		for _, task := range tasks {
			status := "PENDIENTE"
			if task.IsCompleted {
				status = "COMPLETADO"
			}
			fmt.Printf("%d. %s - %s - %s\n", task.ID, task.Title, task.Description, status)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
