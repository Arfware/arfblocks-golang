package todo_queries_all

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

func (m Mapper) mapToResponseModel(todos []entities.Todo) *ResponseModel {

	responseTodos := make([]ResponseModel_Todo, 0)

	for _, todo := range todos {
		responseTodos = append(responseTodos, ResponseModel_Todo{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
		})
	}

	return &ResponseModel{
		Array: responseTodos,
	}
}
