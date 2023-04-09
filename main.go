package main

import (
	"github.com/gin-gonic/gin"
	commonComponent "github.com/rhosocial/go-rush-common/component"
	"github.com/rhosocial/go-rush-manager/controllers/app"
	"log"
)

func main() {
	r := gin.New()
	if !configEngine(r) {
		return
	}
	if err := r.Run(":38080"); err != nil {
		log.Println(err.Error())
		return
	}
}

func configEngine(r *gin.Engine) bool {
	r.Use(
		commonComponent.AppendRequestID(),
		gin.LoggerWithFormatter(commonComponent.LogFormatter),
		commonComponent.AuthRequired(),
		gin.Recovery(),
		commonComponent.ErrorHandler(),
	)
	var ca controllerApp.ControllerApp
	ca.RegisterActions(r)
	return true
}
