package controller

import (
	"github.com/gin-gonic/gin"

	"erpimg/application/common/baseresponse"
	"erpimg/application/imageserver/logic"
)

type (
	ImageController struct {
		imageLogic *logic.ImageLogic
	}
)

func New(imageLogic *logic.ImageLogic) *ImageController {

	return &ImageController{imageLogic: imageLogic}
}

func (c *ImageController) Register(context *gin.Context) {
	r := new(logic.RegisterRequest)
	if err := context.ShouldBindJSON(r); err != nil {
		baseresponse.ParamError(context, err)
		return
	}
	res, err := c.imageLogic.Register(r)
	baseresponse.HttpResponse(context, res, err)
	return
}
