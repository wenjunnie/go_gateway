package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"admin" validate:"required"`
}

func (params *AdminLoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
