package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func GeneralExerciseRoute(router *gin.Engine) {
	router.POST("/generalExercise", controllers.CreateGeneralExercise())
	router.GET("/generalExercise/:generalExerciseId", controllers.GetAGeneralExercise())
	router.PUT("/generalExercise/:generalExerciseId", controllers.EditAGeneralExercise())
	router.DELETE("/generalExercise/:generalExerciseId", controllers.DeleteAGeneralExercise())
	router.GET("/generalExercises", controllers.GetAllGeneralExercises())
}
