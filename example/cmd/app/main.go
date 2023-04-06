package main

import (
	"fmt"

	api "github.com/arfware/arfblocks-golang/example/internal/api/_api"
)

func main() {

	fmt.Println("Waiting for requests...")

	api.Run(7001)
}
