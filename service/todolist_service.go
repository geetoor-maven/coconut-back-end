package service

import (
	"context"
	"go-restful/dto"


)

type TodoListService interface {
	CreatedAtTodoList(ctx context.Context, todoListRequest dto.TodoListRequestDTO) dto.TodoListResponseDTO
	UpdateTodoList(ctx context.Context, todoListRequest dto.UpdateTodoListRequestDTO) dto.TodoListResponseDTO
}