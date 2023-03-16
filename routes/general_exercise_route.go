package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func GeneralExerciseRoute(router *gin.Engine) {
	router.POST("/generalExercise", utils.AuthMiddleware, controllers.CreateGeneralExercise())
	router.GET("/generalExercise/:generalExerciseId", utils.AuthMiddleware, controllers.GetAGeneralExercise())
	router.PUT("/generalExercise/:generalExerciseId", utils.AuthMiddleware, controllers.EditAGeneralExercise())
	router.DELETE("/generalExercise/:generalExerciseId", utils.AuthMiddleware, controllers.DeleteAGeneralExercise())
	router.GET("/generalExercises", utils.AuthMiddleware, controllers.GetAllGeneralExercises())
	router.GET("/generalExercise/join/:generalExerciseId", utils.AuthMiddleware, controllers.GetAGeneralExerciseAndDetails())
}
