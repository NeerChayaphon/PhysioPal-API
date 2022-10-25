package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func AppointmentRoute(router *gin.Engine) {
	router.POST("/appointment", controllers.CreateAppointment())
	router.GET("/appointment/:appointmentId", controllers.GetAAppointment())
	router.PUT("/appointment/:appointmentId", controllers.EditAAppointment())
	router.DELETE("/appointment/:appointmentId", controllers.DeleteAAppointment())
	router.GET("/appointments", controllers.GetAllAppointments())
}
