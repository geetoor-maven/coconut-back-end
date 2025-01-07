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

func (t *todoListRepositoryImpl) UpdateTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) (model.MstTodoList, error) {
	query := `UPDATE mst_todolist 
							SET title = $1, description = $2, status = $3, updated_at = $4
							WHERE id = $5`

	_, err := tx.ExecContext(ctx, query, todoList.Title, todoList.Description, todoList.Status, todoList.UpdatedAt, todoList.ID)
	util.SendPanicIfError(err)

	return todoList, nil
}

func (t *todoListRepositoryImpl) DeleteTodoList(ctx context.Context, tx *sql.Tx, todoList model.MstTodoList) (model.MstTodoList, error) {
	query := `DELETE FROM mst_todolist 
							WHERE id = $1`

	_, err := tx.ExecContext(ctx, query, todoList.ID)
	util.SendPanicIfError(err)

	return todoList, nil
}


func (t *todoListRepositoryImpl) FindTodoListByID(ctx context.Context, tx *sql.Tx, id string) (model.MstTodoList, error) {
	query := "SELECT id, title, description, status, created_at, updated_at FROM mst_todolist WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var todoList model.MstTodoList
	err := row.Scan(&todoList.ID, &todoList.Title, &todoList.Description, &todoList.Status, &todoList.CreatedAt, &todoList.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.MstTodoList{}, nil // ID tidak ditemukan
		}
		return model.MstTodoList{}, err // Error lainnya
	}

	return todoList, nil
}