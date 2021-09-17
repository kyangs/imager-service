package main

import (
	"erpimg/application/common/config"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"erpimg/application/imageserver/controller"
	"erpimg/application/imageserver/logic"
)

var (
	DefaultEngine   = gin.Default()
	DefaultPort     = "9999"
	DefaultFilePath = "/tmp/goodsIdForImageReRsync.list"
)

func main() {
	filePath := flag.String("f", DefaultFilePath, "please use -f /tmp/goodsIdForImageReRsync.list")
	port := flag.String("p", DefaultPort, "please use -p 9999")
	flag.Parse()
	imgCnf := &config.ImageConfig{
		FilePath: *filePath,
		Port:     *port,
	}
	imageController := controller.New(logic.New(imgCnf.FilePath))
	imageRouterGroup := DefaultEngine.Group("/image")
	{
		imageRouterGroup.POST("/register", imageController.Register)
	}
	fmt.Printf("listen at port :%s\n", *port)
	log.Fatal(DefaultEngine.Run(":" + imgCnf.Port))
}
