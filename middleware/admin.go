package middleware

import (
	"final-project-backend/utils/constant"
	"final-project-backend/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	roleId := c.GetInt("roleId")
	if roleId != constant.AdminRoleId {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code":    errors.ErrCodeForbidden,
			"message": errors.ErrForbidden.Error(),
		})
		return
	}
}
