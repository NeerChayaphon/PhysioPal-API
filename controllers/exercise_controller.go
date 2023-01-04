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

var exerciseCollection *mongo.Collection = configs.GetCollection(configs.DB, "exercises")

func CreateExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var exercise models.Exercise
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&exercise); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&exercise); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newExercise := models.Exercise{
			Id:                  primitive.NewObjectID(),
			Details:             exercise.Details,
			MusculoskeltalTypes: exercise.MusculoskeltalTypes,
			Steps:               exercise.Steps,
		}

		result, err := exerciseCollection.InsertOne(ctx, newExercise)
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

func GetAExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		exerciseId := c.Param("exerciseId")
		var exercise models.Exercise
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(exerciseId)

		err := exerciseCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&exercise)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": exercise}})
	}
}

func EditAExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		exerciseId := c.Param("exerciseId")
		var exercise models.Exercise
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(exerciseId)

		//validate the request body
		if err := c.BindJSON(&exercise); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&exercise); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"details": exercise.Details, "musculoskeltalTypes": exercise.MusculoskeltalTypes, "steps": exercise.Steps}
		result, err := exerciseCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated exercise details
		var updatedExercise models.Exercise
		if result.MatchedCount == 1 {
			err := exerciseCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedExercise)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedExercise}})
	}
}

func DeleteAExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		exerciseId := c.Param("exerciseId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(exerciseId)

		result, err := exerciseCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Exercise with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Exercise successfully deleted!"}},
		)
	}
}

func GetAllExercises() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var exercises []models.Exercise
		defer cancel()

		results, err := exerciseCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleExercise models.Exercise
			if err = results.Decode(&singleExercise); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			exercises = append(exercises, singleExercise)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": exercises}},
		)
	}
}
