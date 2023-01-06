package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func TherapeuticExerciseRoutes(router *gin.Engine) {
	router.POST("/therapeuticExercise", utils.AuthMiddleware, controllers.CreateTherapeuticExercise())
	router.GET("/therapeuticExercise/:therapeuticExerciseId", utils.AuthMiddleware, controllers.GetATherapeuticExercise())
	router.PUT("/therapeuticExercise/:therapeuticExerciseId", utils.AuthMiddleware, controllers.EditATherapeuticExercise())
	router.DELETE("/therapeuticExercise/:therapeuticExerciseId", utils.AuthMiddleware, controllers.DeleteATherapeuticExercise())
	router.GET("/therapeuticExercises", utils.AuthMiddleware, controllers.GetAllTherapeuticExercises())
}
