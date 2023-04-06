package todo_queries_all

import (
	"fmt"

	router "github.com/arfware/arfblocks-golang/example/internal/router"
)

type Endpoint struct {
	DataAccess     DataAccess
	Mapper         Mapper
	RequestPayload RequestModel
}

func Init() {
	fmt.Println("Init all")
	requestPayload := RequestModel{}

	endpoint := Endpoint{
		RequestPayload: requestPayload,
	}

	router.Endpoints["/todos/all"] = endpoint
}
