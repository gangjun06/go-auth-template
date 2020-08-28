package auth

import (
	"github.com/gangjun06/book-server/db"
	req "github.com/gangjun06/book-server/models/req"
	resmodel "github.com/gangjun06/book-server/models/res"
	"github.com/gangjun06/book-server/utils"
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func ResetPasswordGetCode(c *gin.Context) {
	body := c.MustGet("body").(*req.ResetPasswordGetCode)

	code := utils.CreateRandomString(8)
	if err := db.SetPasswordVerifyCode(body.Email, code); err != nil {
		if err == db.ErrItemNotFound {
			res.SendError(c, res.ERR_AUTH, "이메일에 해당하는 사용자를 찾을 수 없습니다")
		} else {
			res.SendError(c, res.ERR_SERVER, "서버쪽 에러가 발생하였습니다")
		}
		return
	}
	if err := utils.SendPasswordReset(code, body.Email); err != nil {
		res.SendError(c, res.ERR_SERVER, "인증메일을 전송하는중 에러가 발생하였습니다.")
		return
	}

	res.Response(c, &resmodel.Empty{})
	return
}

func ResetPasswordWithCode(c *gin.Context) {
	body := c.MustGet("body").(*req.ResetPasswordWithCode)
	verified, err := db.CheckPasswordVerify(body.Email, body.Code)
	if !verified {
		if err == db.ErrItemNotFound {
			res.SendError(c, res.ERR_AUTH, "이메일에 해당하는 사용자를 찾을 수 없습니다")
		} else if err != nil {
			res.SendError(c, res.ERR_SERVER, "서버쪽 에러가 발생하였습니다")
		} else {
			res.SendError(c, res.ERR_AUTH, "인증 코드가 일치하지 않습니다")
		}
		return
	}

	if err := db.SetPasswordVerifyCode(body.Email, body.Password); err != nil {
		res.SendError(c, res.ERR_SERVER, "서버쪽 에러가 발생하였습니다")
		return
	}

	res.Response(c, &resmodel.Empty{})
	return
}
