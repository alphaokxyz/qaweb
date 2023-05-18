package database

import (
	"qaweb/errormessage"

	"gorm.io/gorm"
)

type Answer struct {
	Question Question `gorm:"foreignkey:Aid"`
	gorm.Model
	Aid          int    `gorm:"type:int;not null" json:"aid"`
	Answerdetail string `gorm:"type:varchar(1000);not null" json:"answerdetail"`
	Username     string `gorm:"type:varchar(255);not null" json:"username"`
}

func Createanswer(data *Answer) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Getanswersofaquestion(id int, pageSize int, pageNum int) ([]Answer, int) {
	var answers []Answer
	var total int64
	err := db.Preload("Question").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("aid =?", id).Find(&answers).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return answers, int(total)
}

func Deleteanswer(id int) int {
	var answer Answer
	err = db.Where("id = ? ", id).Delete(&answer).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Editanswer(id int, data *Answer) int {
	var answer Answer
	var Answermap = make(map[string]interface{})
	Answermap["aid"] = data.Aid
	Answermap["answerdetail"] = data.Answerdetail
	Answermap["username"] = data.Username

	err = db.Model(&answer).Where("id = ? ", id).Updates(Answermap).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

func Getanswerinfo(id int) (Answer, int) {
	var answer Answer
	err := db.Where("id = ?", id).First(&answer).Error
	if err != nil {
		return answer, errormessage.ERROR_ANSWER_NOT_EXIST
	}
	return answer, errormessage.SUCCESS
}
