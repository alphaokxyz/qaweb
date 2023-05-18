package api

import (
	"net/http"
	"qaweb/database"
	"qaweb/errormessage"
	"qaweb/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

func Adduser(c *gin.Context) {
	var data database.User
	_ = c.ShouldBindJSON(&data)
	var msg string
	msg, code = validator.Validate(&data)
	if code != errormessage.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = database.Checkuser(data.Username)
	if code == errormessage.SUCCESS {
		database.Createuser(&data)
	}
	if code == errormessage.ERROR_USERNAME_USED {
		code = errormessage.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormessage.Geterrormessage(code),
	})
}
func Getusers(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))

	if pagesize == 0 {
		pagesize = -1
	}
	if pagenum == 0 {
		pagenum = 2
	}
	data, total := database.Getusers(pagesize, pagenum)
	code = errormessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormessage.Geterrormessage(code),
	})
}
func Edituser(c *gin.Context) {
	var data database.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = database.Checkuser(data.Username)
	if code == errormessage.SUCCESS {
		database.Edituser(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
	})
}
func Deleteuser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = database.Deleteuser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
	})

}
