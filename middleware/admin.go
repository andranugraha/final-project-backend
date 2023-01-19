package middleware

import (
	"final-project-backend/utils/constant"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	roleId := c.GetInt("roleId")
	if roleId != constant.AdminRoleId {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
}
