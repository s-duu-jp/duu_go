package main

import (
	sw "api/controllers/restapi"
	"api/graph"
	mysql "api/services/db"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	// RESTful API
	router := sw.NewRouter(sw.ApiHandleFunctions{})

	client, _, err := mysql.GetDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Mounting GraphQL and Playground
	router.Any("/query", gin.WrapH(handler.NewDefaultServer(graph.NewSchema(client))))
	router.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	log.Fatal(router.Run(":3000"))
}
