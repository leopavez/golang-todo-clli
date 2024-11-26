package ports

import "todo-cli/internal/domain"

type TaskRepository interface {
	Add(task *domain.Task) error
	GetAll() ([]domain.Task, error)
	CompleteTask(id int) error
}
