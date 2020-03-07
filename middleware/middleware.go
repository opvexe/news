package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"shumin-project/admin-blog-web/model"
	"shumin-project/admin-blog-web/utils"
)

// jwt
func JWT(ctx echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set(echo.HeaderServer, "Echo/999")
		tokenString := context.FormValue("token")
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("Goodsense@"), nil
		})
		if err == nil && token.Valid {
			//验证通过
			context.Set("uid", claims.Id)
			return ctx(context)
		} else {
			return context.JSON(utils.ErrJwt("token验证失败"))
		}
	}
}
