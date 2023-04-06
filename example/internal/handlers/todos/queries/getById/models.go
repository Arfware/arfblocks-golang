package todo_queries_getById

import (
	"github.com/arfware/arfblocks-golang"
)

type RequestModel struct {
	arfblocks.IRequestModel
	Id int `json:"id"`
}

type ResponseModel struct {
	arfblocks.IResponseModel
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
