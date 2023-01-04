package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func TherapeuticExerciseRoutes(router *gin.Engine) {
	router.POST("/therapeuticExercise", controllers.CreateTherapeuticExercise())
	router.GET("/therapeuticExercise/:therapeuticExerciseId", controllers.GetATherapeuticExercise())
	router.PUT("/therapeuticExercise/:therapeuticExerciseId", controllers.EditATherapeuticExercise())
	router.DELETE("/therapeuticExercise/:therapeuticExerciseId", controllers.DeleteATherapeuticExercise())
	router.GET("/therapeuticExercises", controllers.GetAllTherapeuticExercises())
}
