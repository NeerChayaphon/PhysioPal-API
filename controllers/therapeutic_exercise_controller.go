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

var therapeuticExerciseCollection *mongo.Collection = configs.GetCollection(configs.DB, "therapeuticExercises")

func CreateTherapeuticExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var therapeuticExercise models.TherapeuticExercise
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&therapeuticExercise); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&therapeuticExercise); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newTherapeuticExercise := models.TherapeuticExercise{
			Id:            primitive.NewObjectID(),
			AppointmentId: therapeuticExercise.AppointmentId,
			Details:       therapeuticExercise.Details,
			StartDate:     therapeuticExercise.StartDate,
			EndDate:       therapeuticExercise.EndDate,
			ExerciseSet:   therapeuticExercise.ExerciseSet,
		}

		result, err := therapeuticExerciseCollection.InsertOne(ctx, newTherapeuticExercise)
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

func GetATherapeuticExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		therapeuticExerciseId := c.Param("therapeuticExerciseId")
		var therapeuticExercise models.TherapeuticExercise
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(therapeuticExerciseId)

		err := therapeuticExerciseCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&therapeuticExercise)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": therapeuticExercise}})
	}
}

func EditATherapeuticExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		therapeuticExerciseId := c.Param("therapeuticExerciseId")
		var therapeuticExercise models.TherapeuticExercise
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(therapeuticExerciseId)

		//validate the request body
		if err := c.BindJSON(&therapeuticExercise); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&therapeuticExercise); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"details": therapeuticExercise.Details, "startDate": therapeuticExercise.StartDate, "endDate": therapeuticExercise.EndDate, "exerciseSet": therapeuticExercise.ExerciseSet}
		result, err := therapeuticExerciseCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated therapeuticExercise details
		var updatedTherapeuticExercise models.TherapeuticExercise
		if result.MatchedCount == 1 {
			err := therapeuticExerciseCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTherapeuticExercise)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTherapeuticExercise}})
	}
}

func DeleteATherapeuticExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		therapeuticExerciseId := c.Param("therapeuticExerciseId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(therapeuticExerciseId)

		result, err := therapeuticExerciseCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "TherapeuticExercise with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "TherapeuticExercise successfully deleted!"}},
		)
	}
}

func GetAllTherapeuticExercises() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var therapeuticExercises []models.TherapeuticExercise
		defer cancel()

		results, err := therapeuticExerciseCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTherapeuticExercise models.TherapeuticExercise
			if err = results.Decode(&singleTherapeuticExercise); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			therapeuticExercises = append(therapeuticExercises, singleTherapeuticExercise)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": therapeuticExercises}},
		)
	}
}
