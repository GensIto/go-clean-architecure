package usecase

import (
	"go-crean-aarchitecture/model"
	"go-crean-aarchitecture/repository"
)

type ITaskUsecase interface {
	CreateTask(title string) (int, error)
	GetTask(id int) (model.Task, error)
	UpdateTask(id int, title string) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr: tr}
}

func (tu *taskUsecase) CreateTask(title string) (int, error) {
	task := model.Task{Title: title}

	err := task.Validate()
	if err != nil {
		return -1, err
	}

	id, err := tu.tr.Create(&task)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (tu *taskUsecase) GetTask(id int) (model.Task, error) {
	task, err := tu.tr.Read(id)
	if err != nil {
		return model.Task{}, err
	}
	return model.Task{ID: task.ID, Title: task.Title}, nil
}

func (tu *taskUsecase) UpdateTask(id int, title string) error {
	task := model.Task{ID: id, Title: title}
	err := tu.tr.Update(&task)
	if err != nil {
		return err
	}
	return nil
}

func (tu *taskUsecase) DeleteTask(id int) error {
	err := tu.tr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
