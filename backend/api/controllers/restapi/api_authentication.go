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

type AuthenticationAPI struct {
	handlers handlers.AuthenticationHandlers
}

func NewAuthenticationAPI(handlers handlers.AuthenticationHandlers) AuthenticationAPI {
	return AuthenticationAPI{handlers: handlers}
}

// Post /login
// ログイン 
func (api *AuthenticationAPI) PostLogin(c *gin.Context) {
	api.handlers.PostLogin(c)
}
type AuthenticationAPI struct {
	handlers handlers.AuthenticationHandlers
}

func NewAuthenticationAPI(handlers handlers.AuthenticationHandlers) AuthenticationAPI {
	return AuthenticationAPI{handlers: handlers}
}

// Post /logout
// ログアウト 
func (api *AuthenticationAPI) PostLogout(c *gin.Context) {
	api.handlers.PostLogout(c)
}
