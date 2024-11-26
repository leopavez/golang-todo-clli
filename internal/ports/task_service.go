package ports

import "todo-cli/internal/domain"

type TaskService interface {
	AddTask(title, description string) error
	GetAllTask() ([]domain.Task, error)
	CompleteTask(id int) error
}
