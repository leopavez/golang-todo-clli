package repository

import (
	"database/sql"
	"todo-cli/internal/domain"
	"todo-cli/internal/ports"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) ports.TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Add(task *domain.Task) error {
	_, err := r.db.Exec(`INSERT INTO tasks (title, description, isCompleted, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)`, task.Title, task.Description, task.IsCompleted, task.CreatedAt, task.UpdatedAt)
	return err
}

func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, description, isCompleted from tasks`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []domain.Task

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) CompleteTask(id int) error {
	_, err := r.db.Exec(`UPDATE tasks SET isCompleted = true WHERE id = ?`, id)
	return err
}
