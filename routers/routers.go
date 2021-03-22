package routers

import (
	"code.kawai.com/gin/cs/controller"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.LoadHTMLGlob("html/*")
	r.POST("/login", controller.PostLogin)
	r.GET("/login", controller.GetLogin)
	r.GET("/1", controller.HandleLogin)
	r.GET("/choose", controller.GetCourse)
	r.POST("/choose", controller.ChooseCourse)
	return r
}
