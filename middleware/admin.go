package middleware

import "github.com/gin-gonic/gin"

func Admin(c *gin.Context) {
	const adminRoleId = 1
	roleId := c.GetInt("roleId")
	if roleId != adminRoleId {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
}
