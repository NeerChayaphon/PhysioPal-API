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

var musculoskeltalTypesCollection *mongo.Collection = configs.GetCollection(configs.DB, "musculoskeltalTypes")

func CreateMusculoskeltalType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var musculoskeltalType models.MusculoskeltalTypes
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&musculoskeltalType); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&musculoskeltalType); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
			return
		}

		newMusculoskeltalType := models.MusculoskeltalTypes{
			Id:   primitive.NewObjectID(),
			Type: musculoskeltalType.Type,
		}

		result, err := musculoskeltalTypesCollection.InsertOne(ctx, newMusculoskeltalType)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.APIResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}

func GetAMusculoskeltalType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		musculoskeltalTypeId := c.Param("musculoskeltalTypeId")
		var musculoskeltalType models.MusculoskeltalTypes
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(musculoskeltalTypeId)

		err := musculoskeltalTypesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&musculoskeltalType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: musculoskeltalType})
	}
}

func EditAMusculoskeltalType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		musculoskeltalTypeId := c.Param("musculoskeltalTypeId")
		var musculoskeltalType models.MusculoskeltalTypes
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(musculoskeltalTypeId)

		//validate the request body
		if err := c.BindJSON(&musculoskeltalType); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&musculoskeltalType); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
			return
		}

		update := bson.M{"type": musculoskeltalType.Type}
		result, err := musculoskeltalTypesCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//get updated musculoskeltalType details
		var updatedMusculoskeltalType models.MusculoskeltalTypes
		if result.MatchedCount == 1 {
			err := musculoskeltalTypesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedMusculoskeltalType)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: updatedMusculoskeltalType})
	}
}

func DeleteAMusculoskeltalType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		musculoskeltalTypeId := c.Param("musculoskeltalTypeId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(musculoskeltalTypeId)

		result, err := musculoskeltalTypesCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: "MusculoskeltalType with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: "MusculoskeltalType successfully deleted!"},
		)
	}
}

func GetAllMusculoskeltalTypes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var musculoskeltalTypes []models.MusculoskeltalTypes
		defer cancel()

		results, err := musculoskeltalTypesCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleMusculoskeltalType models.MusculoskeltalTypes
			if err = results.Decode(&singleMusculoskeltalType); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			}

			musculoskeltalTypes = append(musculoskeltalTypes, singleMusculoskeltalType)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: musculoskeltalTypes},
		)
	}
}
