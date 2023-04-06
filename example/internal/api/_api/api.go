package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arfware/arfblocks-golang"

	todos_queries_all "github.com/arfware/arfblocks-golang/example/internal/handlers/todos/queries/all"
	todos_queries_getById "github.com/arfware/arfblocks-golang/example/internal/handlers/todos/queries/getById"
	router "github.com/arfware/arfblocks-golang/example/internal/router"
)

func Run(port int16) {

	// SET
	// TODO_CQ
	todos_queries_getById.Init()
	todos_queries_all.Init()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		log.Println(r.Method + " " + r.URL.String())

		endpoint := router.Endpoints[r.URL.Path]
		if endpoint == nil {
			fmt.Println("Not Hit....")
			w.Write([]byte(`{ "message":"NOT FOUND..."}`))
			return
		}
		arfblocks.OperateHttpRequest(w, endpoint)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
