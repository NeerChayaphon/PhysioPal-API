package routes

import (
	"github.com/NeerChayaphon/PhysioPal-API/controllers"
	"github.com/gin-gonic/gin"
)

func MstRoute(router *gin.Engine) {
	router.POST("/musculoskeltalType", controllers.CreateMusculoskeltalType())
	router.GET("/musculoskeltalType/:musculoskeltalTypeId", controllers.GetAMusculoskeltalType())
	router.PUT("/musculoskeltalType/:musculoskeltalTypeId", controllers.EditAMusculoskeltalType())
	router.DELETE("/musculoskeltalType/:musculoskeltalTypeId", controllers.DeleteAMusculoskeltalType())
	router.GET("/musculoskeltalTypes", controllers.GetAllMusculoskeltalTypes())
}
