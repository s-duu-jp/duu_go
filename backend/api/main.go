package main

import (
	sw "api/controllers/restapi"
	"api/handlers/impl"
	"log"
)

func main() {
	// RESTful API
	hs := impl.NewHandlers()
	apiHandleFunctions := sw.ApiHandleFunctions{
		AuthenticationAPI: sw.NewAuthenticationAPI(hs),
		SampleAPI:         sw.NewSampleAPI(hs),
	}
	router := sw.NewRouter(apiHandleFunctions)
	log.Fatal(router.Run(":3000"))
}
