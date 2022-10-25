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

var telemedicineCollection *mongo.Collection = configs.GetCollection(configs.DB, "telemedicines")

func CreateTelemedicine() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var telemedicine models.Telemedicine
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&telemedicine); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&telemedicine); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newTelemedicine := models.Telemedicine{
			Id:              primitive.NewObjectID(),
			RoomID:          telemedicine.RoomID,
			Patient:         telemedicine.Patient,
			Physiotherapist: telemedicine.Physiotherapist,
			Date:            telemedicine.Date,
		}

		result, err := telemedicineCollection.InsertOne(ctx, newTelemedicine)
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

func GetATelemedicine() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		telemedicineId := c.Param("telemedicineId")
		var telemedicine models.Telemedicine
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(telemedicineId)

		err := telemedicineCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&telemedicine)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": telemedicine}})
	}
}

func EditATelemedicine() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		telemedicineId := c.Param("telemedicineId")
		var telemedicine models.Telemedicine
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(telemedicineId)

		//validate the request body
		if err := c.BindJSON(&telemedicine); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&telemedicine); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"roomID": telemedicine.RoomID, "patient": telemedicine.Patient, "physiotherapist": telemedicine.Physiotherapist, "date": telemedicine.Date}
		result, err := telemedicineCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated telemedicine details
		var updatedTelemedicine models.Telemedicine
		if result.MatchedCount == 1 {
			err := telemedicineCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTelemedicine)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTelemedicine}})
	}
}

func DeleteATelemedicine() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		telemedicineId := c.Param("telemedicineId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(telemedicineId)

		result, err := telemedicineCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Telemedicine with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Telemedicine successfully deleted!"}},
		)
	}
}

func GetAllTelemedicines() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var telemedicines []models.Telemedicine
		defer cancel()

		results, err := telemedicineCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTelemedicine models.Telemedicine
			if err = results.Decode(&singleTelemedicine); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			telemedicines = append(telemedicines, singleTelemedicine)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": telemedicines}},
		)
	}
}
