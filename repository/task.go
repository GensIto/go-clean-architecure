package repository

import (
	"database/sql"
	"go-crean-aarchitecture/model"
)

type ITaskRepository interface {
	Create(task *model.Task) (int, error)
	Read(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{db: db}
}

func (tr *taskRepositoryImpl) Create(task *model.Task) (int, error) {
	cmd := `INSERT INTO tasks (title) VALUES (?) RETURNING id`
	err := tr.db.QueryRow(cmd, task.Title).Scan(&task.ID)
	return task.ID, err
}

func (tr *taskRepositoryImpl) Read(id int) (*model.Task, error) {
	cmd := `SELECT id, title FROM tasks WHERE id = ?`
	task := model.Task{}
	err := tr.db.QueryRow(cmd, id).Scan(&task.ID, &task.Title)
	return &task, err
}

func (tr *taskRepositoryImpl) Update(task *model.Task) error {
	cmd := `UPDATE tasks SET title = ? WHERE id = ?`
	rows, err := tr.db.Exec(cmd, task.Title, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (tr *taskRepositoryImpl) Delete(id int) error {
	cmd := `DELETE FROM tasks WHERE id = ?`
	rows, err := tr.db.Exec(cmd, id)
	if err != nil {
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}
