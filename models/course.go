package models

import (
	"errors"

	"code.kawai.com/gin/cs/dao"
)

type CourseStudent struct {
	Sid int `gorm:"sid"`
	Cid int `gorm:"cid"`
}
type CourseTeacher struct {
	Tid     int `gorm:"tid"`
	Cid     int `gorm:"cid"`
	Surplus int `gorm:"surplus"`
}
type Course struct {
	Id   int    `gorm:"id"`
	Cno  int    `gorm:"cno"`
	Name string `gorm:"name"`
}

var (
	s  Student
	t  Teacher
	c  Course
	ct CourseTeacher
	cs CourseStudent
)

func ReduceCourse(course, teacher string, student int) (err error) {
	//fmt.Println(teacher, course)
	if count, err := dao.RDB.Decr(ctx, course+teacher).Result(); err != nil || count < 0 {
		return errors.New("没有课余量了╮(╯-╰)╭")
	}
	dao.DB.Where("name=?", course).Find(&c)
	dao.DB.Where("name=?", teacher).Find(&t)
	dao.DB.Where("cid=? AND tid=?", c.Cno, t.ID).Find(&ct)
	//fmt.Println(teacher, course)
	if ct.Surplus <= 0 {
		return errors.New("没有课余量了╮(╯-╰)╭")
	} else {
		cs.Cid = c.Cno
		cs.Sid = student
		dao.DB.Model(&ct).Where("cid=? AND tid=?", c.Cno, t.ID).Update("surplus", ct.Surplus-1)
		dao.DB.Save(&cs)
	}
	return
}
func GetallCourse() (courstList []*CourseTeacher, err error) {
	if err = dao.DB.Find(&courstList).Error; err != nil {
		return nil, err
	}
	return
}
