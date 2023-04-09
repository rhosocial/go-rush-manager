package controllerApp

import (
	"github.com/gin-gonic/gin"
	commonComponent "github.com/rhosocial/go-rush-common/component"
	controllerAppActivity "github.com/rhosocial/go-rush-manager/controllers/app/activity"
)

type ControllerApp struct {
	commonComponent.GenericController
}

func (c *ControllerApp) RegisterActions(g *gin.Engine) {
	controller := g.Group("/app")
	{
		var ca controllerAppActivity.ControllerAppActivity
		ca.RegisterActions(controller)
	}
}
