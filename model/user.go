package model
import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	UserName string
	Passwordhash string
}


func GetUser(ID interface{}) (User, error) {
	var user User
	result:= DB.First(&user,ID)
	return user, result.Error
}

func (user *User) Setpassword(password string) error {
	 bytes,err :=bcrypt.GenerateFromPassword([]byte(password),12)
	 if err !=nil{
		 return err
	 }
	 user.Passwordhash = string(bytes)
	 return nil
}
func (user *User) Checkpass(password string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(user.Passwordhash), []byte(password))
	return err == nil
}

