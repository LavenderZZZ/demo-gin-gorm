package service

import (
	"demo1/model"
	"demo1/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserRegisterService struct {
	Username string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}
type UserLoginService struct {
	Username string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

func (service *UserLoginService) setSession(c *gin.Context, user model.User)  {
	s :=sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

func (service *UserRegisterService) valid() *serialize.Response {

	count := int64(0)
	model.DB.Model(&model.User{}).Where("user_name = ?", service.Username).Count(&count)
	if count>0 {
		return &serialize.Response{
			Code: 40001,
			Msg: "用户名已经被注册",
		}
	}
	return nil
}

func (service *UserLoginService) Login(c *gin.Context) serialize.Response {

	var user model.User

	if err:=model.DB.Where("user_name = ?", service.Username).First(&user).Error; err!=nil{
		return serialize.ParamErr("账号错误",nil)
	}

	if user.Checkpass(service.Password) ==false{
		return serialize.ParamErr("密码错误",nil)
	}

	service.setSession(c, user)

	return serialize.BuildUserResponse(user)
}

func (service *UserRegisterService) Register() serialize.Response {
	user:=model.User{
		UserName: service.Username,
	}

	if err := user.Setpassword(service.Password) ;err!= nil {
		return serialize.Err(
			serialize.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	if err:= model.DB.Create(&user).Error; err!=nil{
		return serialize.ParamErr("注册失败", err)
	}

	return serialize.BuildUserResponse(user)
}