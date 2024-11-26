package usecase

import (
	"time"
	"todo-cli/internal/domain"
	"todo-cli/internal/ports"
)

type AddTask struct {
	repo ports.TaskRepository
}

func NewAddTask(repo ports.TaskRepository) *AddTask {
	return &AddTask{repo: repo}
}

func (u *AddTask) Execute(title, description string) error {
	task := &domain.Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return u.repo.Add(task)
}
