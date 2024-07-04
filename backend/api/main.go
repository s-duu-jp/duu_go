package main

import (
	sw "api/controllers/restapi"
	"api/handlers/impl"
	"log"
)

func main() {
	// RESTful API
	router := sw.NewRouter(sw.ApiHandleFunctions{
		AuthenticationAPI: sw.NewAuthenticationAPI(impl.NewAuthenticationHandlers()),
	})
	log.Fatal(router.Run(":3000"))
}
