package controller

import (
	"github.com/gin-gonic/gin"
	"go_gateway/dto"
	"go_gateway/middleware"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
