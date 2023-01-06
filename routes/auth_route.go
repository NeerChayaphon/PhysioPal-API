package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

// login rounte for patient
func LoginRoutes(router *gin.Engine) {
	router.POST("/patient/login", utils.PatientLogin())
	router.POST("/physiotherapist/login", utils.PhysiotherapistLogin())
}
