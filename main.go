package main

import (
	"code.kawai.com/gin/cs/dao"
	"code.kawai.com/gin/cs/models"
	"code.kawai.com/gin/cs/routers"
)

func main() {
	r := routers.SetupRouter()
	dao.InitMySQL()
	dao.InitRedis()
	models.CreateAllModel()
	defer dao.Close()
	r.Run(":8000")
}
