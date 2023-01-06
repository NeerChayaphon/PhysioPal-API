package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func PatientRoute(router *gin.Engine) {
	router.POST("/patient", controllers.CreatePatient())
	router.GET("/patient/:patientId", utils.AuthMiddleware, controllers.GetAPatient())
	router.PUT("/patient/:patientId", utils.AuthMiddleware, controllers.EditAPatient())
	router.DELETE("/patient/:patientId", utils.AuthMiddleware, controllers.DeleteAPatient())
	router.GET("/patients", utils.AuthMiddleware, controllers.GetAllPatients())
	router.POST("/patient/exerciseHistory/:patientId", utils.AuthMiddleware, controllers.AddExerciseHistory())
	router.GET("/patient/exerciseHistory/:patientId", utils.AuthMiddleware, controllers.GetPatientExerciseHistory())
}
