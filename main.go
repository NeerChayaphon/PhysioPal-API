package main

import (
	"github.com/NeerChayaphon/PhysioPal-API/configs"
	"github.com/NeerChayaphon/PhysioPal-API/routes"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//run redis
	client := redis.NewClient(&redis.Options{
		Addr:     configs.EnvRedisURI(),
		Password: configs.EnvRedisPassword(),
	})

	//middleware
	router.Use(cors.Default())
	router.Use(utils.CacheMiddleware(client))
	//routes
	routes.LoginRoutes(router)
	routes.PatientRoute(router)
	routes.PhysiotherapistRoute(router)
	routes.ExerciseRoute(router)
	routes.GeneralExerciseRoute(router)
	routes.AppointmentRoute(router)
	routes.TelemedicineRoute(router)
	routes.MstRoute(router)
	routes.TherapeuticExerciseRoutes(router)

	router.Run()
}
