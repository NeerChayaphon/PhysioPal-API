package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
)

func MstRoute(router *gin.Engine) {
	router.POST("/musculoskeltalType", utils.AuthMiddleware, controllers.CreateMusculoskeltalType())
	router.GET("/musculoskeltalType/:musculoskeltalTypeId", utils.AuthMiddleware, controllers.GetAMusculoskeltalType())
	router.PUT("/musculoskeltalType/:musculoskeltalTypeId", utils.AuthMiddleware, controllers.EditAMusculoskeltalType())
	router.DELETE("/musculoskeltalType/:musculoskeltalTypeId", utils.AuthMiddleware, controllers.DeleteAMusculoskeltalType())
	router.GET("/musculoskeltalTypes", utils.AuthMiddleware, controllers.GetAllMusculoskeltalTypes())
}
