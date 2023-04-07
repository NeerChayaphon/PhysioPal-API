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

var patientCollection *mongo.Collection = configs.GetCollection(configs.DB, "patients")

func CreatePatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patient models.Patient
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&patient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
			return
		}

		// find if the email is already registered
		var existingPatient models.Patient
		err := patientCollection.FindOne(ctx, bson.M{"email": patient.Email}).Decode(&existingPatient)
		if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "email already registered"})
			return
		}

		passwordHash, err := utils.HashPassword(patient.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		newPatient := models.Patient{
			Id:                primitive.NewObjectID(),
			Name:              patient.Name,
			Email:             patient.Email,
			Password:          passwordHash,
			Phone:             patient.Phone,
			Photo:             patient.Photo,
			Address:           patient.Address,
			CongenitalDisease: patient.CongenitalDisease,
			ExerciseHistory:   make([]models.ExerciseHistory, 0),
		}

		result, err := patientCollection.InsertOne(ctx, newPatient)
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

func GetAPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId := c.Param("patientId")
		var patient models.Patient
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(patientId)

		err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: patient})
	}
}

func EditAPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId := c.Param("patientId")
		var patient models.Patient
		var patientTemp models.Patient
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(patientId)

		err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patientTemp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//validate the request body
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&patient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
			return
		}

		var passwordHash string

		if patient.Password != "" {
			passwordHash, err = utils.HashPassword(patient.Password)
		} else {
			passwordHash = patientTemp.Password
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		update := bson.M{"name": patient.Name, "email": patient.Email, "password": passwordHash, "phone": patient.Phone, "photo": patient.Photo, "address": patient.Address, "congenitalDisease": patient.CongenitalDisease, "gender": patient.Gender}
		result, err := patientCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//get updated patient details
		var updatedPatient models.Patient
		if result.MatchedCount == 1 {
			err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPatient)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: updatedPatient})
	}
}

func DeleteAPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId := c.Param("patientId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(patientId)

		result, err := patientCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{Status: http.StatusNotFound, Message: "error", Data: "Patient with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: "Patient successfully deleted!"},
		)
	}
}

func GetAllPatients() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patients []models.Patient
		defer cancel()

		results, err := patientCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singlePatient models.Patient
			if err = results.Decode(&singlePatient); err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			}

			patients = append(patients, singlePatient)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{Status: http.StatusOK, Message: "success", Data: patients},
		)
	}
}

func AddExerciseHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId := c.Param("patientId")
		var exerciseHistory models.ExerciseHistory
		var patient models.Patient
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(patientId)

		//validate the request body
		if err := c.BindJSON(&exerciseHistory); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&exerciseHistory); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
			return
		}

		// Continue here
		newExerciseHistory := models.ExerciseHistory{
			Date:             time.Now(),
			ExerciseType:     exerciseHistory.ExerciseType,
			ExerciseSetId:    exerciseHistory.ExerciseSetId,
			ExerciseRecorded: exerciseHistory.ExerciseRecorded,
			IsComplete:       exerciseHistory.IsComplete,
			ExerciseStatus:   exerciseHistory.ExerciseStatus,
		}

		err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		// update patient Exercise History
		patient.ExerciseHistory = append(patient.ExerciseHistory, newExerciseHistory)

		update := patient
		result, err := patientCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//get updated patient details
		var updatedPatient models.Patient
		if result.MatchedCount == 1 {
			err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPatient)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: updatedPatient})
	}
}

func GetPatientExerciseHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId := c.Param("patientId")
		var patient models.Patient
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(patientId)

		err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: patient.ExerciseHistory})
	}
}
