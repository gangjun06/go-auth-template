package auth

import (
	"github.com/gangjun06/book-server/db"
	req "github.com/gangjun06/book-server/models/req"
	resmodel "github.com/gangjun06/book-server/models/res"
	"github.com/gangjun06/book-server/utils"
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	body := c.MustGet("body").(*req.SignIn)
	result, err := db.FindUserByEmail(body.Email)
	if err != nil {
		if err == db.ErrItemNotFound {
			res.SendError(c, res.ERR_AUTH, "유저정보를 찾을 수 없습니다")
		} else {
			res.SendError(c, res.ERR_SERVER, "SERVER ERROR")
		}
		return
	}

	if result := utils.CheckPassword(body.Password, result.Password); result == false {
		res.SendError(c, res.ERR_AUTH, "비밀번호가 일치하지 않습니다")
		return
	}

	token, err := utils.GetJwtToken(result.ID)
	if err != nil {
		res.SendError(c, res.ERR_SERVER, "SERVER_ERROR ERROR")
		return
	}

	res.Response(c, resmodel.Auth{
		Token:    token,
		Verified: result.Verified,
		Name:     result.Name,
		Email:    result.Email,
	})
}
