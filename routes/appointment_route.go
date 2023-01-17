package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func AppointmentRoute(router *gin.Engine) {
	router.POST("/appointment", utils.AuthMiddleware, controllers.CreateAppointment())
	router.GET("/appointment/:appointmentId", utils.AuthMiddleware, controllers.GetAAppointment())
	router.PUT("/appointment/:appointmentId", utils.AuthMiddleware, controllers.EditAAppointment())
	router.DELETE("/appointment/:appointmentId", utils.AuthMiddleware, controllers.DeleteAAppointment())
	router.GET("/appointments", utils.AuthMiddleware, controllers.GetAllAppointments())

	router.GET("/appointments/patient/:patientId", utils.AuthMiddleware, controllers.GetAppointmentsByPatient())
	router.GET("/appointments/physiotherapist/:physiotherapistId", utils.AuthMiddleware, controllers.GetAppointmentsByPhysiotherapist())
}
