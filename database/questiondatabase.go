package database

import (
	"qaweb/errormessage"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Title          string `gorm:"type:varchar(1000);not null" json:"title"`
	Questiondetail string `gorm:"type:varchar(1000);not null" json:"questiondetail"`
	Username       string `gorm:"type:varchar(255);not null" json:"username"`
}

func Createquestion(data *Question) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Getquestions(pagesize int, pagenum int) ([]Question, int) {
	var questions []Question
	var total int64

	err := db.Order("id DESC").Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&questions).Error
	if err != nil {
		return nil, 0
	}

	err = db.Model(&Question{}).Count(&total).Error
	if err != nil {
		return nil, 0
	}

	return questions, int(total)
}

func Deletequestion(id int) int {
	var question Question
	err = db.Where("id = ? ", id).Delete(&question).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Editquestion(id int, data *Question) int {
	var question Question
	var Questionmap = make(map[string]interface{})
	Questionmap["title"] = data.Title
	Questionmap["questiondetail"] = data.Questiondetail
	Questionmap["username"] = data.Username

	err = db.Model(&question).Where("id = ? ", id).Updates(Questionmap).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Getquestioninfo(id int) (Question, int) {
	var question Question
	err := db.Where("id = ?", id).First(&question).Error
	if err != nil {
		return question, errormessage.ERROR_QUESTION_NOT_EXIST
	}
	return question, errormessage.SUCCESS
}
