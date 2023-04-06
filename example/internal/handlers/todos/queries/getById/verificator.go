package todo_queries_getById

import "errors"

func (e Endpoint) Verificate() error {
	dataIndex := e.DataAccess.GetTodoIndexById(e.RequestPayload.Id)
	if dataIndex == -1 {
		return errors.New("todo not exist")
	}

	return nil
}
