package main

import (
	"github.com/NeerChayaphon/PhysioPal-API/configs"
	"github.com/NeerChayaphon/PhysioPal-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
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
