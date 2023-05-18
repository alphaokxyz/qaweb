package api

import (
	"net/http"
	"qaweb/database"
	"qaweb/errormessage"
	"qaweb/middleware"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Addanswer(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	splitResult := strings.Split(tokenString, " ")
	tokenstring := splitResult[1]
	username := middleware.Getusername(tokenstring)
	var data database.Answer
	_ = c.ShouldBindJSON(&data)
	data.Username = username

	database.Createanswer(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormessage.Geterrormessage(code),
	})
}
func Getanswersofaquestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))

	if pagesize == 0 {
		pagesize = -1
	}
	if pagenum == 0 {
		pagenum = 2
	}
	data, total := database.Getanswersofaquestion(id, pagesize, pagenum)
	code = errormessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormessage.Geterrormessage(code),
	})
}
func Editanswer(c *gin.Context) {
	var data database.Answer
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	database.Editanswer(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
	})
}
func Deleteanswer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = database.Deleteanswer(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
	})

}

func Getanswerinfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := database.Getanswerinfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormessage.Geterrormessage(code),
	})
}
