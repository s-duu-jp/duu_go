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
}

// Post /login
// ログイン 
func (api *AuthenticationAPI) PostLogin(c *gin.Context) {
	factory := func() interface{} {
		return &PostLoginRequest{}
	}
	handlers.PostLogin(c, factory)
}

// Post /logout
// ログアウト 
func (api *AuthenticationAPI) PostLogout(c *gin.Context) {
	factory := func() interface{} {
		return &PostLogoutRequest{}
	}
	handlers.PostLogout(c, factory)
}

// Post /refresh-token
// トークン再発行 
func (api *AuthenticationAPI) PostRefreshToken(c *gin.Context) {
	factory := func() interface{} {
		return &PostRefreshTokenRequest{}
	}
	handlers.PostRefreshToken(c, factory)
}

