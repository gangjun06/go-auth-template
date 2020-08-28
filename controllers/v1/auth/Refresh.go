package auth

import (
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gangjun06/book-server/db"
	"github.com/gangjun06/book-server/models"
	resmodel "github.com/gangjun06/book-server/models/res"
	"github.com/gangjun06/book-server/utils"
	"github.com/gangjun06/book-server/utils/res"
	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	// Parsing Token From Header
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		res.SendError(c, res.ERR_BAD_REQUEST, "토큰을 넣어주세요")
		return
	}

	extractedToken := strings.Split(clientToken, "Bearer ")

	// Verify if the format of the token is correct
	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		res.SendError(c, res.ERR_BAD_REQUEST, "토큰을 올바른 포멧으로 입력해주세요")
		return
	}

	// Parsing JWT To struct
	claims := &models.Claims{}
	_, errParseWithClaims := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetConfig().JwtSecret), nil
	})

	// Check Correct OR Has Error
	if errParseWithClaims != nil {
		log.Println(errParseWithClaims)
		if errParseWithClaims.Error() == jwt.ErrSignatureInvalid.Error() {
			res.SendError(c, res.ERR_BAD_REQUEST, "올바른 토큰 포멧이 아닙니다")
		} else {
			token, err := utils.GetJwtToken(claims.ID)
			if err != nil {
				res.SendError(c, res.ERR_SERVER, "SERVER_ERROR ERROR")
				return
			}
			res.Response(c, &resmodel.Token{
				Token: token,
			})
		}
		return
	}

	_, err := db.FindUserByID(claims.ID)
	if err != nil {
		if err == db.ErrItemNotFound {
			res.SendError(c, res.ERR_AUTH, "유저정보를 찾을 수 없습니다")
		} else {
			res.SendError(c, res.ERR_SERVER, "SERVER ERROR")
		}
		return
	}

	res.Response(c, &resmodel.Token{
		Token: clientToken,
	})
}
