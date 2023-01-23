package middleware

import (
	"final-project-backend/utils/constant"
	"final-project-backend/utils/errors"
	"final-project-backend/utils/response"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	roleId := c.GetInt("roleId")
	if roleId != constant.UserRoleId {
		response.SendError(c, 403, errors.ErrCodeForbidden, errors.ErrForbidden.Error())
		return
	}
}
