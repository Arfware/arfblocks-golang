package todo_queries_all

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

func (d DataAccess) GetAll() []entities.Todo {
	return mockTodos
}
