package usecase

import (
	"todo-cli/internal/domain"
	"todo-cli/internal/ports"
)

type GetAllTask struct {
	repo ports.TaskRepository
}

func NewGetAllTask(repo ports.TaskRepository) *GetAllTask {
	return &GetAllTask{repo: repo}
}

func (u *GetAllTask) Execute() ([]domain.Task, error) {
	return u.repo.GetAll()
}
