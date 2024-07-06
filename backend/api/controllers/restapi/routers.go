/*
 * プロダクト名：Sample
 *
 * APIの説明
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

type  struct {
	handlers handlers.Handlers
}

func New(handlers handlers.Handlers)  {
	return {handlers: handlers}
}

```

1. `/workspace/openapi/config/go/templates/routers.mustache`編集

```go
/*
 * プロダクト名：Sample
 *
 * APIの説明
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"api/config/env"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	router := gin.Default()

	// Load configuration
	cfg, err := env.GetConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// 環境変数をチェックして本番環境以外の場合にCORSミドルウェアを追加
	if env := cfg["ENV"]; env != "production" {
		router.Use(cors.Default())
	}

	return NewRouterWithGinEngine(router, handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the AuthenticationAPI part of the API
	AuthenticationAPI AuthenticationAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"PostLogin",
			http.MethodPost,
			"/login",
			handleFunctions.AuthenticationAPI.PostLogin,
		},
		{
			"PostLogout",
			http.MethodPost,
			"/logout",
			handleFunctions.AuthenticationAPI.PostLogout,
		},
		{
			"PostRefreshToken",
			http.MethodPost,
			"/refresh-token",
			handleFunctions.AuthenticationAPI.PostRefreshToken,
		},
	}
}