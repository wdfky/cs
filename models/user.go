package models

import "code.kawai.com/gin/cs/dao"

type UserInfo struct {
	Sno      string `gorm:"sno"`
	Password string `gorm:"password"`
}
type Student struct {
	ID       int    `gorm:"id"`
	ClazzId  int    `gorm:"spassword"`
	Gender   int    `gorm:"sname"`
	Name     string `gorm:"ssex"`
	Notes    string `gorm:"sage"`
	Password string `gorm:"password"`
	Sno      string `gorm:"sno"`
}
type Teacher struct {
	ID   int    `gorm:"id"`
	Name string `gorm:"name"`
}

func CheckUser(u *Student) bool {
	var stu Student
	dao.DB.Where("sno=?", u.Sno).Find(&stu)
	if u.Password == stu.Password {
		return true
	}
	return false
}
