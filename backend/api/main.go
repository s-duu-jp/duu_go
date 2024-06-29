package main

import (
	sw "api/controllers/restapi"
	"api/graph"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	// RESTful API
	router := sw.NewRouter(sw.ApiHandleFunctions{})

	// Mounting GraphQL and Playground
	gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Any("/query", gin.WrapH(gqlServer))
	router.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	log.Fatal(router.Run(":3000"))
}
