package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/NeerChayaphon/PhysioPal-API/configs"
	"github.com/NeerChayaphon/PhysioPal-API/models"
	"github.com/NeerChayaphon/PhysioPal-API/responses"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var patientCollection *mongo.Collection = configs.GetCollection(configs.DB, "patients")
var physiotherapistCollection *mongo.Collection = configs.GetCollection(configs.DB, "physiotherapists")

func PatientLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate the login credentials
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patient models.Patient
		var loginDetails LoginDetails
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&loginDetails); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		// find the patient by email and password
		err := patientCollection.FindOne(ctx, bson.M{"email": loginDetails.Email}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusUnauthorized, Message: "error", Data: "invalid email or password"})
			return
		}

		if !CheckPasswordHash(loginDetails.Password, patient.Password) {
			c.JSON(http.StatusUnauthorized, responses.APIResponse{Status: http.StatusUnauthorized, Message: "error", Data: "invalid email or password"})
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = patient.Id // change later
		claims["role"] = "patient"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": t,
			"role":  "patient",
			"id":    patient.Id,
		})
	}
}

func PhysiotherapistLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate the login credentials
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var physiotherapist models.Physiotherapist
		var loginDetails LoginDetails
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&loginDetails); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		// find the physiotherapist by email
		err := physiotherapistCollection.FindOne(ctx, bson.M{"email": loginDetails.Email}).Decode(&physiotherapist)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusUnauthorized, Message: "error", Data: "invalid email or password"})
			return
		}

		if !CheckPasswordHash(loginDetails.Password, physiotherapist.Password) {
			c.JSON(http.StatusUnauthorized, responses.APIResponse{Status: http.StatusUnauthorized, Message: "error", Data: "invalid email or password"})
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = physiotherapist.Id // change later
		claims["role"] = "physiotherapist"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"token":  t,
			"role":   "physiotherapist",
			"id":     physiotherapist.Id,
		})
	}
}

func AuthMiddleware(c *gin.Context) {
	// Extract the JWT from the request header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Need JWT Token",
		})
		return
	}
	// Validate the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// If the JWT is valid, set the user ID and role in the context
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.Set("user_id", claims["id"])
	c.Set("user_role", claims["role"])

	c.Next()
}

func GetUserByJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the user ID and role from the context
		userId := c.MustGet("user_id").(string)
		userRole := c.MustGet("user_role").(string)

		if userRole == "patient" {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var patient models.Patient
			defer cancel()

			objId, _ := primitive.ObjectIDFromHex(userId)
			err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"role":    userRole,
				"data":    patient,
			})
		}

		if userRole == "physiotherapist" {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var physiotherapist models.Physiotherapist
			defer cancel()

			objId, _ := primitive.ObjectIDFromHex(userId)
			err := physiotherapistCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&physiotherapist)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			}

			// c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: { role: user_role physiotherapist}})
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"role":    userRole,
				"data":    physiotherapist,
			})
		}

	}
}
