package main

import (
	sw "api/controllers/restapi"
	"api/handlers/impl"
	"log"
)

func main() {
	// RESTful API
	apiHandleFunctions := sw.ApiHandleFunctions{
		AuthenticationAPI: sw.NewAuthenticationAPI(
			impl.NewAuthenticationHandlers(),
		),
	}
	router := sw.NewRouter(apiHandleFunctions)
	log.Fatal(router.Run(":3000"))
}
