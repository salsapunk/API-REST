package repository

import (
	"database/sql"
	"fmt"

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
	// SELECT nas rows
	rows, err := tr.connection.Query(models.ShowTasksSQL)
	if err != nil {
		fmt.Println(err)
		return []models.Task{}, err
	}

	// cria variaveis para iteração
	var taskList []models.Task
	var taskObj models.Task

	// iteração que adiciona os valores das rows no db dentro do taskObj
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

		// adiciona o taskObj no taskList
		taskList = append(taskList, taskObj)
	}

	// fecha as rows
	err = rows.Close()
	if err != nil {
		fmt.Println(err)
		return []models.Task{}, err
	}

	// retorna a lista com todas as tasks
	return taskList, nil
}

func (tr *TaskRepository) Create(task *models.Task) (int, error) {
	var id int

	// cria a query com o Prepare()
	query, err := tr.connection.Prepare(models.CreateTaskDSQL)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// QueryRow() executa a query e retorna uma *Row
	// Scan() copia as colunas da row nos valores apontados pelo destino
	// err assume o valor de error dado pela função Scan()
	err = query.QueryRow(
		task.Title,
		task.Description,
		task.Done,
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
