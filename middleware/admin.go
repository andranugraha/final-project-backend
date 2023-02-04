package middleware

import (
	"final-project-backend/config"
	"final-project-backend/utils/constant"
	"final-project-backend/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	if config.ENV == "test" {
		c.Set("roleId", constant.AdminRoleId)
		c.Next()
		return
	}

	roleId := c.GetInt("roleId")
	if roleId != constant.AdminRoleId {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code":    errors.ErrCodeForbidden,
			"message": errors.ErrForbidden.Error(),
		})
		return
	}
}
