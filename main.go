package main

import (
	"code.kawai.com/gin/cs/dao"
	"code.kawai.com/gin/cs/routers"
)

func main() {
	r := routers.SetupRouter()
	dao.InitMySQL()
	dao.InitRedis()
	defer dao.Close()
	r.Run(":8000")
}
