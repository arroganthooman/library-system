package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header["Authorization"]
		if len(header) <= 0 {
			ctx.JSON(http.StatusForbidden, "Need Bearer token")
			ctx.Abort()
			return
		}
		token := strings.Split(header[0], "Bearer ")
		if len(token) < 2 {
			ctx.JSON(http.StatusForbidden, "Need bearer token")
			ctx.Abort()
			return
		}

		username, err := m.repositories.CheckToken(token[1])
		if err != nil {
			ctx.JSON(http.StatusForbidden, "Bearer token invalid")
			ctx.Abort()
			return
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}
