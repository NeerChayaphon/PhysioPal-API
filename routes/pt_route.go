package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func PhysiotherapistRoute(router *gin.Engine) {
	router.POST("/physiotherapist", controllers.CreatePhysiotherapist())
	router.GET("/physiotherapist/:physiotherapistId", controllers.GetAPhysiotherapist())
	router.PUT("/physiotherapist/:physiotherapistId", controllers.EditAPhysiotherapist())
	router.DELETE("/physiotherapist/:physiotherapistId", controllers.DeleteAPhysiotherapist())
	router.GET("/physiotherapists", controllers.GetAllPhysiotherapists())
}
