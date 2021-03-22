package models

import (
	"context"
	"time"

	"code.kawai.com/gin/cs/dao"
)

var ctx = context.Background()

func CreateAllModel() {
	courseList, _ := GetallCourse()
	var c Course
	var t Teacher
	for _, course := range courseList {
		dao.DB.Where("cno=?", course.Cid).Find(&c)
		dao.DB.Where("id=?", course.Tid).Find(&t)
		dao.RDB.Set(ctx, c.Name+t.Name, course.Surplus, 48*time.Hour)
	}
}
