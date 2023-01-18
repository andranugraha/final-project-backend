package response

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func SendSuccessWithPagination(c *gin.Context, statusCode int, data interface{}, totalRows int64, totalPages int) {
	c.JSON(statusCode, gin.H{
		"data":       data,
		"totalRows":  totalRows,
		"totalPages": totalPages,
	})
}

func SendError(c *gin.Context, statusCode int, errCode string, msg string) {
	c.JSON(statusCode, gin.H{
		"code":    errCode,
		"message": msg,
	})
}
