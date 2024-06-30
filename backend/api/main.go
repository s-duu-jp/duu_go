package main

import (
	sw "api/controllers/restapi"
	"log"

	mysql "api/services/db"

	"api/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	// RESTful API
	router := sw.NewRouter(sw.ApiHandleFunctions{})

	// Mounting GraphQL and Playground
	// シングルトン化されたデータベースクライアントを取得
	client, _, err := mysql.GetDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Close()

	//gqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{client}}))

	srv := handler.NewDefaultServer(resolver.NewSchema(client))

	router.Any("/query", gin.WrapH(srv))
	router.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	log.Fatal(router.Run(":3000"))
}
