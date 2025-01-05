package repository

import (
	"context"
	"database/sql"
	"go-restful/model"
	"go-restful/util"
)

type todoListRepositoryImpl struct {
}

func NewTodoListRepositoryImpl() TodoListRepository {
	return &todoListRepositoryImpl{}
}

func (t *todoListRepositoryImpl) CreateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) (model.MstTodoList, error) {

	query := `INSERT INTO mst_todolist 
		(id, title, description, status, created_at)
		VALUES 
		($1, $2, $3, $4, $5)`

	_, err := tx.ExecContext(ctx, query, todoList.ID, todoList.Title, todoList.Description, todoList.Status, todoList.CreatedAt)
	util.SendPanicIfError(err)

	return todoList, nil
}

func (t *todoListRepositoryImpl) UpdateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList)  (model.MstTodoList, error){
	query := `UPDATE mst_todoList SET title = $1, description = $2, status = $3, update_at = $4 WHERE id = $5`

	_, err := tx.ExecContext(ctx, query, todoList.Title, todoList.Description, todoList.Status, todoList.UpdatedAt, todoList.ID)
	util.SendPanicIfError(err)

	return todoList, nil
}