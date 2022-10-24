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

var generalExerciseCollection *mongo.Collection = configs.GetCollection(configs.DB, "generalExercises")

func CreateGeneralExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var generalExercise models.GeneralExercise
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&generalExercise); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&generalExercise); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newGeneralExercise := models.GeneralExercise{
			Name:                generalExercise.Name,
			Description:         generalExercise.Description,
			MusculoskeltalTypes: generalExercise.MusculoskeltalTypes,
			Functional:          generalExercise.Functional,
			ExerciseSet:         generalExercise.ExerciseSet,
		}

		result, err := generalExerciseCollection.InsertOne(ctx, newGeneralExercise)
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

func GetAGeneralExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		generalExerciseId := c.Param("generalExerciseId")
		var generalExercise models.GeneralExercise
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(generalExerciseId)

		err := generalExerciseCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&generalExercise)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": generalExercise}})
	}
}

func EditAGeneralExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		generalExerciseId := c.Param("generalExerciseId")
		var generalExercise models.GeneralExercise
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(generalExerciseId)

		//validate the request body
		if err := c.BindJSON(&generalExercise); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&generalExercise); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"name": generalExercise.Name, "description": generalExercise.Description, "musculoskeltalTypes": generalExercise.MusculoskeltalTypes, "functional": generalExercise.Functional, "exerciseSet": generalExercise.ExerciseSet}
		result, err := generalExerciseCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated generalExercise details
		var updatedGeneralExercise models.GeneralExercise
		if result.MatchedCount == 1 {
			err := generalExerciseCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedGeneralExercise)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedGeneralExercise}})
	}
}

func DeleteAGeneralExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		generalExerciseId := c.Param("generalExerciseId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(generalExerciseId)

		result, err := generalExerciseCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "GeneralExercise with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "GeneralExercise successfully deleted!"}},
		)
	}
}

func GetAllGeneralExercises() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var generalExercises []models.GeneralExercise
		defer cancel()

		results, err := generalExerciseCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleGeneralExercise models.GeneralExercise
			if err = results.Decode(&singleGeneralExercise); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			generalExercises = append(generalExercises, singleGeneralExercise)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": generalExercises}},
		)
	}
}
