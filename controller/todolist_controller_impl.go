package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-restful/service"
	"net/http"
	"go-restful/dto"
	"go-restful/util"
)

type todoListControllerImpl struct {
	TodoListService service.TodoListService
}


func NewTodoListControllerImpl(todoListService service.TodoListService) TodoListController {
	return &todoListControllerImpl{
			TodoListService: todoListService,
	}
}

func (t todoListControllerImpl) CreateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestDTO := dto.TodoListRequestDTO{}
	util.ReadFromRequestBody(request, &requestDTO)

	t.TodoListService.CreatedAtTodoList(request.Context(), requestDTO)

	util.WriteToResponseBody(writer, requestDTO)
}

func (t todoListControllerImpl) UpdateTodoList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestDTO := dto.TodoListRequestDTO{}
	util.ReadFromRequestBody(request, &requestDTO)

	t.TodoListService.UpdateTodoList(request.Context(), requestDTO)

	util.WriteToResponseBody(writer, requestDTO)
}
