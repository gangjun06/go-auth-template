package auth

import (
	"net/http"

	"github.com/gangjun06/book-server/db"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	code := c.Param("code")
	check, err := db.CheckVerify(code)
	if err != nil {
		if err == db.ErrUserAlreadyVerified {
			c.HTML(http.StatusOK, "verify.html", gin.H{"text": "이미 인증된 계정입니다.", "success": false})
		} else {

			c.HTML(http.StatusOK, "verify.html", gin.H{"text": "서비스 처리중 에러가 발생하였습니다", "success": false})
		}
		return
	}

	if check {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "정상적으로 메일인증이 되었습니다.\n이제 창을닫고 로그인하실 수 있습니다", "success": true})
	} else {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "URL을 찾지 못하였습니다.\n제대로된 요청인지 확인해주세요", "success": false})
	}
}
