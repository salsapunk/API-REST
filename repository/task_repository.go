package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/salsapunk/API-REST/models"
)

type TaskRepository struct {
	connection *sql.DB
}

func NewTaskRepository(connection *sql.DB) TaskRepository {
	return TaskRepository{
		connection: connection,
	}
}

func (tr *TaskRepository) ListAll() ([]models.Task, error) {
	rows, err := tr.connection.Query(models.ShowTasks)
	if err != nil {
		fmt.Println(err)
		return []models.Task{}, err
	}

	var taskList []models.Task
	var taskObj models.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Title,
			&taskObj.Description,
			&taskObj.Done,
			&taskObj.CreatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []models.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	err = rows.Close()
	if err != nil {
		fmt.Println(err)
		return []models.Task{}, err
	}

	return taskList, nil
}

func (tr *TaskRepository) ListByID(id int) (*models.Task, error) {
	query, err := tr.connection.Prepare(models.ShowTaskBI)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var taskObj models.Task

	err = query.QueryRow(id).Scan(&taskObj.ID,
		&taskObj.Title,
		&taskObj.Description,
		&taskObj.Done,
		&taskObj.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	err = query.Close()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &taskObj, nil
}

func (tr *TaskRepository) Create(task *models.Task) (int, error) {
	var id int

	query, err := tr.connection.Prepare(models.CreateTask)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	task.CreatedAt = time.Now()

	err = query.QueryRow(
		task.Title,
		task.Description,
		task.CreatedAt).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (tr *TaskRepository) Edit(id int) error {
	query, err := tr.connection.Prepare(models.EditTask)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = query.QueryRow(id).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
