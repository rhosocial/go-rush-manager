package controllerAppActivity

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	commonComponent "github.com/rhosocial/go-rush-common/component"
	"net/http"
	"time"
)

type ControllerAppActivity struct {
	commonComponent.GenericController
}

type ActivityBodyCommon struct {
	ClientID uint32 `form:"client_id" validate:"required" json:"client_id" xml:"client_id" binding:"required"`
}

func (c *ControllerAppActivity) RegisterActions(r *gin.RouterGroup) {
	controller := r.Group("/activity")
	{
		controller.PUT("", c.ActionAdd)
	}
}

type ActivityBodyAdd struct {
	ActivityBodyCommon
	EndsAt        time.Time `form:"ends_at" validate:"required" json:"ends_at" xml:"ends_at" binding:"required"`
	Seats         []uint32  `form:"seats" json:"seats" xml:"seats"`
	QueueStartsAt time.Time `form:"queue_starts_at" validate:"required" json:"queue_starts_at" xml:"queue_starts_at" binding:"required"`
	QueueEndsAt   time.Time `form:"queue_ends_at" validate:"required" json:"queue_ends_at" xml:"queue_ends_at" binding:"required"`
}

func (c *ControllerAppActivity) ActionAdd(g *gin.Context) {
	var body ActivityBodyAdd
	if err := g.ShouldBindWith(&body, binding.FormPost); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, commonComponent.NewGenericResponse(g, 1, "request parsing failed", err.Error(), nil))
		return
	}
	g.JSON(http.StatusOK, commonComponent.NewGenericResponse(g, 0, "success", nil, nil))
}
