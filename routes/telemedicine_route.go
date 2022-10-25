package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func TelemedicineRoute(router *gin.Engine) {
	router.POST("/telemedicine", controllers.CreateTelemedicine())
	router.GET("/telemedicine/:telemedicineId", controllers.GetATelemedicine())
	router.PUT("/telemedicine/:telemedicineId", controllers.EditATelemedicine())
	router.DELETE("/telemedicine/:telemedicineId", controllers.DeleteATelemedicine())
	router.GET("/telemedicines", controllers.GetAllTelemedicines())
}
