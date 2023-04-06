package todo_queries_getById

import (
	router "github.com/arfware/arfblocks-golang/example/internal/router"
)

type Endpoint struct {
	DataAccess     DataAccess
	Mapper         Mapper
	RequestPayload RequestModel
}

func Init() {
	requestPayload := RequestModel{
		Id: int(1),
	}

	endpoint := Endpoint{
		RequestPayload: requestPayload,
	}

	router.Endpoints["/todos/get"] = endpoint
}
