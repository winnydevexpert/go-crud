package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"go-crud-learn/config"
	constant "go-crud-learn/constant"
	utils "go-crud-learn/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func BasicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.AppConfig.Mode == constant.MA_MODE_SHORT_TEXT || config.AppConfig.Mode == constant.MA_MODE_FULL_TEXT {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, utils.ErrorMessage(constant.ERROR_SERVICE_UNAVAILABLE, nil))
		} else {
			c.Next()
		}
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.AppConfig.Mode == constant.MA_MODE_SHORT_TEXT || config.AppConfig.Mode == constant.MA_MODE_FULL_TEXT {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, utils.ErrorMessage(constant.ERROR_SERVICE_UNAVAILABLE, nil))
		} else {
			auth := strings.Split(c.GetHeader("Authorization"), " ")
			if len(auth) < 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorMessage(constant.ERROR_UNAUTHORIZED, nil))
				return
			}
			tokenString := auth[1]
			token, errCheckJWT := jwt.ParseWithClaims(tokenString, &utils.JwtToken{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.AppConfig.AccessSecretKey), nil
			})

			if errCheckJWT != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorMessage(constant.ERROR_UNAUTHORIZED, nil))
				return
			}

			if claims, ok := token.Claims.(*utils.JwtToken); ok && token.Valid {
				c.Set("currentUser", claims)
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorMessage(constant.ERROR_UNAUTHORIZED, nil))
			}
		}
	}
}
