package main

import (
	sw "api/controllers/restapi"
	"api/handlers/restapi"
	"log"
)

func main() {
	// RESTful API
	router := sw.NewRouter(sw.ApiHandleFunctions{
		AuthenticationAPI: sw.NewAuthenticationAPI(restapi.NewAuthenticationHandlers()),
	})
	log.Fatal(router.Run(":3000"))
}
