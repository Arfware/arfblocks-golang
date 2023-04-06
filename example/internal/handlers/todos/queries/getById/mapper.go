package todo_queries_getById

import (
	"github.com/arfware/arfblocks-golang/example/internal/entities"
)

type Mapper struct {
}

// func (m Mapper) mapToRequestModel(request *http.Request) *RequestModel {
// 	params := mux.Vars(request)
// 	id, _ := strconv.ParseInt(params["id"], 10, 32)

// 	return &RequestModel{Id: int(id)}
// }

func (m Mapper) mapToResponseModel(todo *entities.Todo) *ResponseModel {

	return &ResponseModel{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
	}
}
