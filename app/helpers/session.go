package helpers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func CheckToken(ctx http.Context) bool {

	token := ctx.Request().Headers().Get("Authorization")

	_, err := facades.Auth().Parse(ctx, token)

	if err != nil {
		return false
	}
	return true

}
