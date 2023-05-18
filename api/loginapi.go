package api

import (
	"net/http"
	"qaweb/database"
	"qaweb/errormessage"
	"qaweb/middleware"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data database.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	code = database.Checkwhenlogin(data.Username, data.Password)

	if code == errormessage.SUCCESS {
		token, code = middleware.Maketoken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
		"token":   token,
	})
}
