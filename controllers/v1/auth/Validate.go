package auth

import (
	resmodel "github.com/gangjun06/book-server/models/res"
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	res.Response(c, resmodel.Empty{})
}
