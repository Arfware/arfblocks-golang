package todoApi

import (
	"fmt"
	"net/http"

	"github.com/arfware/arfblocks-golang"
	todo_queries_all "github.com/arfware/arfblocks-golang/example/internal/handlers/todos/queries/all"
	todo_queries_getById "github.com/arfware/arfblocks-golang/example/internal/handlers/todos/queries/getById"
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/todos/get/{id}", Get).Methods("GET")
	router.HandleFunc("/todos/all", Get).Methods("GET")
}

func Test(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	fmt.Println("TESTSTETSTETSTESTSEETET")
}

func Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	// params := mux.Vars(request)
	// id, _ := strconv.ParseInt(params["id"], 10, 32)

	// Create request payload
	requestPayload := todo_queries_getById.RequestModel{
		Id: int(1),
	}

	// Create endpoint
	endpoint := todo_queries_getById.Endpoint{
		RequestPayload: requestPayload,
	}

	// Run endpoint
	arfblocks.OperateHttpRequest(response, endpoint)
}

func All(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	// Create request payload
	requestPayload := todo_queries_all.RequestModel{}

	// Create endpoint
	endpoint := todo_queries_all.Endpoint{
		RequestPayload: requestPayload,
	}

	// Run endpoint
	arfblocks.OperateHttpRequest(response, endpoint)
}
