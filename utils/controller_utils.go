package utils

import "github.com/gin-gonic/gin"

// Response function will return the response in the user requested format
func Response(c *gin.Context, status int, body interface{}) {
	accType := c.GetHeader("Accept")
	if accType == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func ResponseError(c *gin.Context, status int, errMsg *ApplicationError) {
	accType := c.GetHeader("Accept")
	if accType == "application/xml" {
		c.XML(status, errMsg)
		return
	}
	c.JSON(status, errMsg)
}
