package controllers

import (
	"context"

	"net/http"
	"time"

	"github.com/NeerChayaphon/PhysioPal-API/configs"
	"github.com/NeerChayaphon/PhysioPal-API/models"
	"github.com/NeerChayaphon/PhysioPal-API/responses"
	"github.com/NeerChayaphon/PhysioPal-API/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var appointmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "appointments")

func CreateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var appointment models.Appointment
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&appointment); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Get Patient by ID
		var patient models.Patient
		err := patientCollection.FindOne(ctx, bson.M{"_id": appointment.Patient}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		// Get Physiotherapist by ID
		var physiotherapist models.Physiotherapist
		err = physiotherapistCollection.FindOne(ctx, bson.M{"_id": appointment.Physiotherapist}).Decode(&physiotherapist)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		newAppointment := models.Appointment{
			Id:                  primitive.NewObjectID(),
			Patient:             appointment.Patient,
			Physiotherapist:     appointment.Physiotherapist,
			Date:                appointment.Date,
			Injury:              appointment.Injury,
			Treatment:           appointment.Treatment,
			TherapeuticExercise: appointment.TherapeuticExercise,
		}

		result, err := appointmentCollection.InsertOne(ctx, newAppointment)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.APIResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		appointmentId := c.Param("appointmentId")
		var appointment models.Appointment
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(appointmentId)

		err := appointmentCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&appointment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": appointment}})
	}
}

func EditAAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		appointmentId := c.Param("appointmentId")
		var appointment models.Appointment
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(appointmentId)

		//validate the request body
		if err := c.BindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&appointment); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"patient": appointment.Patient, "physiotherapist": appointment.Physiotherapist, "date": appointment.Date, "injury": appointment.Injury, "treatment": appointment.Treatment, "therapeuticExercise": appointment.TherapeuticExercise}
		result, err := appointmentCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated appointment details
		var updatedAppointment models.Appointment
		if result.MatchedCount == 1 {
			err := appointmentCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAppointment)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAppointment}})
	}
}

func DeleteAAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		appointmentId := c.Param("appointmentId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(appointmentId)

		result, err := appointmentCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Appointment with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Appointment successfully deleted!"}},
		)
	}
}

func GetAllAppointments() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var appointments []models.Appointment
		defer cancel()

		results, err := appointmentCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleAppointment models.Appointment
			if err = results.Decode(&singleAppointment); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			appointments = append(appointments, singleAppointment)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": appointments}},
		)
	}
}

// Create testing for CreateAAppointment function in appointment_test.go
// func TestCreateAAppointment(t *testing.T) {
// 	//setup
// 	gin.SetMode(gin.TestMode)
// 	r := gin.Default()
// 	r.POST("/api/v1/appointment", CreateAAppointment())

// 	//create a new appointment
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/api/v1/appointment", strings.NewReader(`{"patient":"5f8a8d1c2c9d440000c0b7f7","physiotherapist":"5f8a8d1c2c9d440000c0b7f7","date":"2020-10-19T00:00:00Z","injury":"5f8a8d1c2c9d440000c0b7f7","treatment":"5f8a8d1c2c9d440000c0b7f7","therapeuticExercise":"5f8a8d1c2c9d440000c0b7f7"}`))
// 	r.ServeHTTP(w, req)

// 	//assertions
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
// 	assert.Equal(t, `{"status":200,"message":"success","data":{"data":{"id":"5f8a8d1c2c9d440000c0b7f7","patient":"5f8a8d1c2c9d440000c0b7f7","physiotherapist":"5f8a8d1c2c9d440000c0b7f7","date":"2020-10-19T00:00:00Z","injury":"5f8a8d1c2c9d440000c0b7f7","treatment":"5f8a8d1c2c9d440000c0b7f7","therapeuticExercise":"5f8a8d1c2c9d440000c0b7f7"}}}`, w.Body.String())
// }
