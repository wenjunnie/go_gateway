package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
	"time"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"admin" validate:"required"`
}

func (params *AdminLoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` //token
}
