package todo_queries_all

func (e Endpoint) Handle() (any, error) {
	todos := e.DataAccess.GetAll()
	mappedResponse := e.Mapper.mapToResponseModel(todos)

	return mappedResponse, nil
}
