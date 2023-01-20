package middleware

import (
	"final-project-backend/utils/constant"
	"final-project-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	roleId := c.GetInt("roleId")
	if roleId != constant.AdminRoleId {
		response.SendError(c, 403, "FORBIDDEN", "You are not authorized to access this resource")
		return
	}
}
