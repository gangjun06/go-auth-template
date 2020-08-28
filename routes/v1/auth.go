package v1

import (
	c "github.com/gangjun06/book-server/controllers/v1/auth"
	m "github.com/gangjun06/book-server/middlewares"
	req "github.com/gangjun06/book-server/models/req"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(r *gin.RouterGroup) {
	r.GET("/verify/:code", c.Verify)
	r.POST("/signup", m.VerifyRequest(&req.SignUp{}), c.SignUp)
	r.POST("/signin", m.VerifyRequest(&req.SignIn{}), c.SignIn)
	r.GET("/validate", m.CheckAuth(), c.Validate)
	r.GET("/refresh", c.Refresh)
	r.POST("/updatepassword", m.VerifyRequest(&req.ResetPassword{}), m.CheckAuth(), c.UpdatePassword)
	r.POST("/resetpasswordcode", m.VerifyRequest(&req.ResetPasswordGetCode{}), c.ResetPasswordGetCode)
	r.POST("/resetpasswordwithcode", m.VerifyRequest(&req.ResetPasswordWithCode{}), c.ResetPasswordWithCode)
}
