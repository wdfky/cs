package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"code.kawai.com/gin/cs/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostLogin(c *gin.Context) {
	var user models.Student
	user.Sno = c.PostForm("sno")
	user.Password = c.PostForm("passward")
	if ok := models.CheckUser(&user); !ok {
		session := sessions.Default(c)
		session.Set("sno", user.Sno)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请检查学号密码",
		})
	}
}
func GetLogin(c *gin.Context) {
	s := c.GetHeader("Cookie")
	if s == "" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		session := sessions.Default(c)
		sno := session.Get("sno")
		if sno != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "sno",
				"sno":     sno,
			})
		} else {
			c.HTML(http.StatusOK, "login.html", nil)
		}
	}
}
func HandleLogin(c *gin.Context) {
	session := sessions.Default(c)
	sno := session.Get("sno")
	if sno != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": sno,
			"sno":     sno,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "tmd",
		})
	}
}
func ChooseCourse(c *gin.Context) {
	teacher := c.PostForm("teacher")
	course := c.PostForm("course")
	//fmt.Println(teacher, course)
	session := sessions.Default(c)
	sno := session.Get("sno")
	// fmt.Println(teacher, course)
	// fmt.Println(sno.(string))
	//这是一个叫str的int类型
	str, err := strconv.Atoi(sno.(string))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", str)
	fmt.Println(str)
	err = models.ReduceCourse(course, teacher, str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "选课失败",
		})
		
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "选课成功",
		})
	}
}
func GetCourse(c *gin.Context) {
	s := c.GetHeader("Cookie")
	if s == "" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		session := sessions.Default(c)
		sno := session.Get("sno")
		if sno != nil {
			c.HTML(http.StatusOK, "choose.html", nil)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "sno",
				"sno":     sno,
			})
		}
	}
}
