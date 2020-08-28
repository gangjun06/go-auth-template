package auth

import (
	"github.com/gangjun06/book-server/db"
	req "github.com/gangjun06/book-server/models/req"
	resmodel "github.com/gangjun06/book-server/models/res"
	"github.com/gangjun06/book-server/utils"
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	body := c.MustGet("body").(*req.SignUp)
	if err := db.CreateUser(body.Email, body.Name, body.Password); err != nil {
		res.SendError(c, res.ERR_DUPLICATE, "해당 이메일 주소는 이미 가입되어 있습니다.")
		return
	}

	result, err := db.FindUserByEmail(body.Email)
	if err == db.ErrItemNotFound {
		res.SendError(c, res.ERR_SERVER, "Error Load User Info")
		return
	}

	token, errGetJwtToken := utils.GetJwtToken(result.ID)
	if err := utils.SendVefiryMail(result.VerifyCode, body.Email); err != nil {
		res.SendError(c, res.ERR_SERVER, "Error Send Verify Mail")
		return
	}
	if errGetJwtToken != nil {
		res.SendError(c, res.ERR_SERVER, "Error Get JWT Token")
		return
	}

	res.Response(c, &resmodel.Auth{
		Token:    token,
		Verified: false,
		Name:     result.Name,
		Email:    result.Email,
	})
}
