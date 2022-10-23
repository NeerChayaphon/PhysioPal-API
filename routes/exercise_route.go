package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func ExerciseRoute(router *gin.Engine) {
	router.POST("/exercise", controllers.CreateExercise())
	router.GET("/exercise/:exerciseId", controllers.GetAExercise())
	router.PUT("/exercise/:exerciseId", controllers.EditAExercise())
	router.DELETE("/exercise/:exerciseId", controllers.DeleteAExercise())
	router.GET("/exercises", controllers.GetAllExercises())
}
