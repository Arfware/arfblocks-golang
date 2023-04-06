package todo_queries_getById

func (e Endpoint) Handle() (any, error) {
	matchedTodo := e.DataAccess.GetTodoById(e.RequestPayload.Id)
	mappedResponse := e.Mapper.mapToResponseModel(matchedTodo)

	return mappedResponse, nil
}
