package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func PhysiotherapistRoute(router *gin.Engine) {
	router.POST("/physiotherapist", utils.AuthMiddleware, controllers.CreatePhysiotherapist())
	router.GET("/physiotherapist/:physiotherapistId", utils.AuthMiddleware, controllers.GetAPhysiotherapist())
	router.PUT("/physiotherapist/:physiotherapistId", utils.AuthMiddleware, controllers.EditAPhysiotherapist())
	router.DELETE("/physiotherapist/:physiotherapistId", utils.AuthMiddleware, controllers.DeleteAPhysiotherapist())
	router.GET("/physiotherapists", utils.AuthMiddleware, controllers.GetAllPhysiotherapists())
}
