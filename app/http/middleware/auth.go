package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/helpers"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		result := helpers.CheckToken(ctx)

		if !result {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Request().Next()
	}
}
