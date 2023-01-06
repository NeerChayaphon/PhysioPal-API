package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func TelemedicineRoute(router *gin.Engine) {
	router.POST("/telemedicine", utils.AuthMiddleware, controllers.CreateTelemedicine())
	router.GET("/telemedicine/:telemedicineId", utils.AuthMiddleware, controllers.GetATelemedicine())
	router.PUT("/telemedicine/:telemedicineId", utils.AuthMiddleware, controllers.EditATelemedicine())
	router.DELETE("/telemedicine/:telemedicineId", utils.AuthMiddleware, controllers.DeleteATelemedicine())
	router.GET("/telemedicines", utils.AuthMiddleware, controllers.GetAllTelemedicines())
}
