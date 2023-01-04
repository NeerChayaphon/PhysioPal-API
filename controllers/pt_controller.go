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

var physiotherapistCollection *mongo.Collection = configs.GetCollection(configs.DB, "physiotherapists")

func CreatePhysiotherapist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var physiotherapist models.Physiotherapist
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&physiotherapist); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&physiotherapist); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		passwordHash, err := utils.HashPassword(physiotherapist.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newPhysiotherapist := models.Physiotherapist{
			Id:       primitive.NewObjectID(),
			Details:  physiotherapist.Details,
			Password: passwordHash,
			Phone:    physiotherapist.Phone,
			Photo:    physiotherapist.Photo,
		}

		result, err := physiotherapistCollection.InsertOne(ctx, newPhysiotherapist)
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

func GetAPhysiotherapist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		physiotherapistId := c.Param("physiotherapistId")
		var physiotherapist models.Physiotherapist
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(physiotherapistId)

		err := physiotherapistCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&physiotherapist)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": physiotherapist}})
	}
}

func EditAPhysiotherapist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		physiotherapistId := c.Param("physiotherapistId")
		var physiotherapist models.Physiotherapist
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(physiotherapistId)

		//validate the request body
		if err := c.BindJSON(&physiotherapist); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&physiotherapist); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		passwordHash, err := utils.HashPassword(physiotherapist.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := bson.M{"details": physiotherapist.Details, "password": passwordHash, "phone": physiotherapist.Phone, "photo": physiotherapist.Photo}
		result, err := physiotherapistCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated physiotherapist details
		var updatedPhysiotherapist models.Physiotherapist
		if result.MatchedCount == 1 {
			err := physiotherapistCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPhysiotherapist)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPhysiotherapist}})
	}
}

func DeleteAPhysiotherapist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		physiotherapistId := c.Param("physiotherapistId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(physiotherapistId)

		result, err := physiotherapistCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Physiotherapist with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Physiotherapist successfully deleted!"}},
		)
	}
}

func GetAllPhysiotherapists() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var physiotherapists []models.Physiotherapist
		defer cancel()

		results, err := physiotherapistCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singlePhysiotherapist models.Physiotherapist
			if err = results.Decode(&singlePhysiotherapist); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			physiotherapists = append(physiotherapists, singlePhysiotherapist)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": physiotherapists}},
		)
	}
}
