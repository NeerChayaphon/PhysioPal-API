package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func PatientRoute(router *gin.Engine) {
	router.POST("/patient", controllers.CreatePatient())
	router.GET("/patient/:patientId", controllers.GetAPatient())
	router.PUT("/patient/:patientId", controllers.EditAPatient())
	router.DELETE("/patient/:patientId", controllers.DeleteAPatient())
	router.GET("/patients", controllers.GetAllPatients())
	router.POST("/patient/exerciseHistory/:patientId", controllers.AddExerciseHistory())
	router.GET("/patient/exerciseHistory/:patientId", controllers.GetPatientExerciseHistory())
}
