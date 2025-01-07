package service

import (
	"context"
	"database/sql"
	"go-restful/dto"
	"go-restful/model"
	"go-restful/repository"
	"go-restful/util"
	"time"

	"github.com/google/uuid"
)

type todoListServiceImpl struct {
	TodoListRepository repository.TodoListRepository

	DB *sql.DB
}

func NewTodoListServiceImpl(todoListRepository repository.TodoListRepository, DB *sql.DB) TodoListService {
	return &todoListServiceImpl{
		TodoListRepository: todoListRepository,
		DB:                DB,
	}
}

func (t *todoListServiceImpl) CreatedAtTodoList(ctx context.Context, todoListRequest dto.TodoListRequestDTO) dto.TodoListResponseDTO {
	tx, err := t.DB.Begin()
	util.SendPanicIfError(err)

	defer util.CommitOrRollback(tx)

	todoList := model.MstTodoList{
		ID:          uuid.New().String(),
		Title:       todoListRequest.Title,
		Description: todoListRequest.Description,
		Status:      todoListRequest.Status,
		CreatedAt:   time.Now(),
	}

	createTodoList, errSave := t.TodoListRepository.CreateTodoList(ctx, tx, todoList)
    util.SendPanicIfError(errSave)

    return convertToResponseDTO(createTodoList)

	
}

func convertToResponseDTO(mstTodoList model.MstTodoList) dto.TodoListResponseDTO {
	return dto.TodoListResponseDTO{
			Title:       mstTodoList.Title,
			Description: mstTodoList.Description,
			Status:      mstTodoList.Status,
			CreatedAt:   time.Now(),
	}
} 

func (t *todoListServiceImpl) UpdateTodoList(ctx context.Context, todoListRequest dto.UpdateTodoListRequestDTO) dto.TodoListResponseDTO {
	tx, err := t.DB.Begin()
	util.SendPanicIfError(err)

	defer util.CommitOrRollback(tx)

	// Cek apakah ID ada di database
	existingTodoList, err := t.TodoListRepository.FindTodoListByID(ctx, tx, todoListRequest.ID)
	util.SendPanicIfError(err)

	// Jika ID tidak ditemukan
	if existingTodoList.ID == "" {
		panic("ID not found")
	}

	// Jika ID ditemukan, lakukan update
	todoList := model.MstTodoList{
		ID:          todoListRequest.ID,
		Title:       todoListRequest.Title,
		Description: todoListRequest.Description,
		Status:      todoListRequest.Status,
		UpdatedAt:   time.Now(),
	}

	updatedTodoList, errSave := t.TodoListRepository.UpdateTodoList(ctx, tx, todoList)
	util.SendPanicIfError(errSave)

	return convertToResponseDTO(updatedTodoList)
}

func (t *todoListServiceImpl) DeleteTodoList(ctx context.Context, todoListRequest dto.DeleteTodoListRequestDTO) dto.TodoListResponseDTO {
	tx, err := t.DB.Begin()
	util.SendPanicIfError(err)

	defer util.CommitOrRollback(tx)

	todoList := model.MstTodoList{
			ID: todoListRequest.ID,
	}

	deletedTodoList, errSave := t.TodoListRepository.DeleteTodoList(ctx, tx, todoList)
	util.SendPanicIfError(errSave)

	return convertToResponseDTO(deletedTodoList)
}