package models

import "code.kawai.com/gin/cs/dao"

type UserInfo struct {
	Sno      string `gorm:"sno"`
	Password string `gorm:"password"`
}
type Student struct {
	Sno       string `gorm:"sno"`
	Spassword string `gorm:"spassword"`
	Sname     string `gorm:"sname"`
	Ssex      string `gorm:"ssex"`
	Sage      int    `gorm:"sage"`
	Sdapt     string `gorm:"sdapt"`
	Sid       string `gorm:"sid"`
	Spower    string `gorm:"spower"`
}

func CheckUser(u *Student) bool {
	var stu Student
	dao.DB.Where("sno=?", u.Sno).Find(&stu)
	if u.Spassword == stu.Spassword {
		return true
	}
	return false
}
