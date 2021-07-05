package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go_gateway/dao"
	"go_gateway/dto"
	"go_gateway/middleware"
	"time"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLoginOut)
}

// 声明一个全局的redisdb变量
var redisdb *redis.Client

// InitClient 初始化连接
func InitClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}

	//1. params.UserName 取得管理员信息 admininfo
	//2. admininfo.salt + params.Password sha256 => saltPassword
	//3. saltPassword==admininfo.password
	tx, err := lib.GetGormPool("default") //default为配置文件中设定的名字
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//在Redis中设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo) //格式化
	if err != nil {
		return
	}
	err = redisdb.Set("session_admin", string(sessBts), -1).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	//sessInfo := &dto.AdminSessionInfo{
	//	ID:        admin.Id,
	//	UserName:  admin.UserName,
	//	LoginTime: time.Now(),
	//}
	//sessBts, err := json.Marshal(sessInfo) //格式化
	//if err != nil {
	//	middleware.ResponseError(c, 2003, err)
	//	return
	//}
	//sess := sessions.Default(c)
	//sess.Set(public.AdminSessionInfoKey, string(sessBts))
	//sess.Save()

	out := dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLoginOut 退出登录
func (adminLogin *AdminLoginController) AdminLoginOut(c *gin.Context) {
	redisdb.Del("session_admin")
	middleware.ResponseSuccess(c, "")
}
