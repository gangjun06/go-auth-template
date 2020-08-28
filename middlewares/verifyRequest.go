package middlewares

import (
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func VerifyRequest(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(data); err != nil {
			res.SendError(c, res.ERR_BAD_REQUEST, err.Error())
			return
		}
		c.Set("body", data)
	}
}
