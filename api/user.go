package api

import (
	"demo1/conf"
	"demo1/model"
	"demo1/serialize"
	service2 "demo1/service"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)
func Ping(c *gin.Context) {
	c.JSON(200, serialize.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := interface{}(user).(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serialize.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serialize.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serialize.ParamErr("JSON类型不匹配", err)
	}

	return serialize.ParamErr("参数错误", err)
}

func UserLogin(c *gin.Context)  {
	var service service2.UserLoginService
	if err:= c.ShouldBind(&service); err==nil{
		res:= service.Login(c)
		c.JSON(200,res)
	}else{
		c.JSON(200, ErrorResponse(err))
	}
}

func UserRegister(c *gin.Context)  {
	var service service2.UserRegisterService
	if err:= c.ShouldBind(&service); err==nil{
		res:= service.Register()
		c.JSON(200,res)
	}else{
		c.JSON(200, ErrorResponse(err))
	}
}

func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serialize.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serialize.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
