package dto

type TodoListRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}

type UpdateTodoListRequestDTO struct {
	ID          string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}
