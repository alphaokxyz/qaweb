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

func Addquestion(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	splitResult := strings.Split(tokenString, " ")
	tokenstring := splitResult[1]
	username := middleware.Getusername(tokenstring)
	var data database.Question
	_ = c.ShouldBindJSON(&data)
	data.Username = username

	database.Createquestion(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormessage.Geterrormessage(code),
	})
}
func Getquestions(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))

	if pagesize == 0 {
		pagesize = -1
	}
	if pagenum == 0 {
		pagenum = 2
	}
	data, total := database.Getquestions(pagesize, pagenum)
	code = errormessage.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormessage.Geterrormessage(code),
	})
}
func Editquestion(c *gin.Context) {
	var data database.Question
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	database.Editquestion(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
	})
}
func Deletequestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = database.Deletequestion(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormessage.Geterrormessage(code),
	})

}

func Getquestioninfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := database.Getquestioninfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormessage.Geterrormessage(code),
	})
}
