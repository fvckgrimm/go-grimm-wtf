package middleware

import (
	"github.com/goravel/framework/contracts/http"
    "strings"
)

func DetectMobile() http.Middleware {
	return func(ctx http.Context) {
		userAgent := ctx.Request().Header("User-Agent")
        isMobile := strings.Contains(strings.ToLower(userAgent), "mobile")
        ctx.WithValue("isMobile", isMobile)
        ctx.Request().Next()
	}
}
