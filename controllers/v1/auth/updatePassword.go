package auth

import (
	"github.com/gangjun06/book-server/db"
	dbmodels "github.com/gangjun06/book-server/models/db"
	reqmodel "github.com/gangjun06/book-server/models/req"
	resmodel "github.com/gangjun06/book-server/models/res"
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func UpdatePassword(c *gin.Context) {
	user := c.MustGet("User").(*dbmodels.User)
	body := c.MustGet("body").(*reqmodel.ResetPassword)
	if err := db.UpdatePassword(user.ID, body.Password); err != nil {
		res.SendError(c, res.ERR_SERVER, "Error Load User Info")
		return
	}
	res.Response(c, &resmodel.Empty{})
}
