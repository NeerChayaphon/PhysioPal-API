package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func ExerciseRoute(router *gin.Engine) {
	router.POST("/exercise", utils.AuthMiddleware, controllers.CreateExercise())
	router.GET("/exercise/:exerciseId", utils.AuthMiddleware, controllers.GetAExercise())
	router.PUT("/exercise/:exerciseId", utils.AuthMiddleware, controllers.EditAExercise())
	router.DELETE("/exercise/:exerciseId", utils.AuthMiddleware, controllers.DeleteAExercise())
	router.GET("/exercises", utils.AuthMiddleware, controllers.GetAllExercises())
}
