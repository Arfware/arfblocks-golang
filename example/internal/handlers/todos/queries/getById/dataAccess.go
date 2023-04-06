package todo_queries_getById

import (
	"github.com/arfware/arfblocks-golang/example/internal/entities"
)

type DataAccess struct {
}

var mockTodos = []entities.Todo{
	{
		Id:          1,
		Title:       "Creating Employee List",
		Description: "Please create a list that contains all employee information",
		CreatedBy:   "Mary Gleen",
	},
	{
		Id:          3,
		Title:       "Printing Advertisement Cards",
		Description: "Please print 1000 advertisement card that sent via email before",
		CreatedBy:   "John Doe",
	}}

func (d DataAccess) GetTodoById(id int) *entities.Todo {
	// Querying Data
	for _, v := range mockTodos {
		if v.Id == id {
			return &v
		}
	}

	return nil
}

func (d DataAccess) GetTodoIndexById(id int) int {
	// Querying Data
	for _, v := range mockTodos {
		if v.Id == id {
			return id
		}
	}

	return -1
}
