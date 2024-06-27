package main

import (
	sw "api/controllers/restapi"
	"log"
)

func main() {
	// RESTful API
	router := sw.NewRouter(sw.ApiHandleFunctions{})
	log.Fatal(router.Run(":3000"))
}