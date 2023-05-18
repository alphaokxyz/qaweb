package database

import (
	"encoding/base64"
	"log"
	"qaweb/errormessage"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

func Checkuser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errormessage.ERROR_USERNAME_USED
	}
	return errormessage.SUCCESS
}

func Createuser(data *User) int {
	data.Password = Scryptpassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Getusers(pagesize int, pagenum int) ([]User, int) {
	var users []User
	var total int64
	err = db.Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&users).Count(&total).Error
	if err != nil {
		return nil, 0
	}

	return users, int(total)
}

func Scryptpassword(password string) string {

	salt := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	dk, err := scrypt.Key([]byte(password), salt, 2, 8, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	scryptpassword := base64.StdEncoding.EncodeToString(dk)
	return scryptpassword
}

func Deleteuser(id int) int {
	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Edituser(id int, data *User) int {
	var user User
	var Usermap = make(map[string]interface{})
	Usermap["username"] = data.Username
	Usermap["role"] = data.Role
	err = db.Model(&user).Where("id = ? ", id).Updates(Usermap).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Checkwhenlogin(username string, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return errormessage.ERROR_USER_NOT_EXIST
	}
	if user.Password != Scryptpassword(password) {
		return errormessage.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 && user.Role != 2 {
		return errormessage.ERROR_USER_NO_RIGHT
	}
	return errormessage.SUCCESS

}
