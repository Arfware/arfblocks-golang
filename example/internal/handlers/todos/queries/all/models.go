package todo_queries_all

import (
	"github.com/arfware/arfblocks-golang"
)

type RequestModel struct {
	arfblocks.IRequestModel
}

type ResponseModel struct {
	arfblocks.IResponseModel
	Array []ResponseModel_Todo `json:"array"`
}

type ResponseModel_Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
