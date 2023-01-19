package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"final-project-backend/config"
	errResp "final-project-backend/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticated(c *gin.Context) {
	tokenString, err := getJwtToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errResp.ErrCodeUnauthorized,
			"message": err.Error(),
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})

	if err != nil {
		log.Println(err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    errResp.ErrCodeUnauthorized,
					"message": "not a token",
				})
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    errResp.ErrCodeUnauthorized,
					"message": "token expired or not actived yet",
				})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    errResp.ErrCodeUnauthorized,
					"message": "couldn't handle token",
				})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errResp.ErrCodeUnauthorized,
			"message": "couldn't handle token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		intId := int(claims["userId"].(float64))
		c.Set("userId", intId)

		intRoleId := int(claims["roleId"].(float64))
		c.Set("roleId", intRoleId)

		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    errResp.ErrCodeUnauthorized,
		"message": "couldn't handle token",
	})
}

func getJwtToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
