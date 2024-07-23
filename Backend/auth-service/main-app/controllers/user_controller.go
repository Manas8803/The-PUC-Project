package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/db"
	network "github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/lib/net"
	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/lib/security"
	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/main-app/models"
	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/main-app/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func Login(r *gin.Context) {

	var req models.Login

	//* Checking for invalid json format
	if err := r.BindJSON(&req); err != nil {
		network.RespondWithError(r, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(&req); validationErr != nil {
		network.RespondWithError(r, http.StatusBadRequest, "Please provide the required credentials.")
		return
	}

	//* Checking whether the user is registered
	user, userErr := db.GetUserByEmail(req.Email)
	if userErr != nil {
		log.Println(userErr)
		if strings.Contains(userErr.Error(), "no authorities registered with the given email address") {
			network.RespondWithError(r, http.StatusNotFound, "Email is not registered with any authorities.")
			return
		}
		network.RespondWithError(r, http.StatusInternalServerError, "Internal server error : "+userErr.Error())
		return
	}

	//* Checking Password
	securityErr := security.CheckPassword(req.Password, user.Password)
	if securityErr != nil {
		network.RespondWithError(r, http.StatusUnauthorized, "Invalid Credentials : Password does not match")
		return
	}

	//* Generating Token
	token, genJWTErr := security.GenerateJWT()
	if genJWTErr != nil {
		network.RespondWithError(r, http.StatusInternalServerError, "Internal Server Error : "+genJWTErr.Error())
		return
	}

	r.JSON(http.StatusOK, responses.UserResponse{Message: "success", Data: map[string]interface{}{"token": "Bearer " + token, "office_name": user.OfficeName}})
}

func Register(r *gin.Context) {
	r.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	r.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var req models.Register

	//* Checking for invalid json format and populating "user"
	if invalidJsonErr := r.BindJSON(&req); invalidJsonErr != nil {
		network.RespondWithError(r, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(&req); validationErr != nil {
		network.RespondWithError(r, http.StatusUnprocessableEntity, "Please provide the required credentials")
		return
	}

	//* Hashing Password
	hashedPass, hashPassErr := security.HashPassword(req.Password)
	if hashPassErr != nil {
		network.RespondWithError(r, http.StatusInternalServerError, "Internal Server Error : "+hashPassErr.Error())
		return
	}

	//* Creating User
	_, insertDBErr := db.CreateUserByEmail(&db.User{
		OfficeName: req.OfficeName,
		Email:      req.Email,
		Password:   hashedPass,
	})

	//* Checking for errors while inserting in the DB
	if insertDBErr != nil {
		log.Println(insertDBErr)
		if strings.Contains(insertDBErr.Error(), "ConditionalCheckFailedException: The conditional request failed") {
			network.RespondWithError(r, http.StatusConflict, "Email is already registered. Please login")
			return
		}

		network.RespondWithError(r, http.StatusInternalServerError, insertDBErr.Error()+"  : Error in inserting the document")
		return
	}

	//* Generating Token
	token, genJWTErr := security.GenerateJWT()
	if genJWTErr != nil {
		network.RespondWithError(r, http.StatusInternalServerError, "Internal Server Error : "+genJWTErr.Error())
		return
	}

	r.JSON(http.StatusCreated, responses.UserResponse{Message: "You email has been successfully registered", Data: map[string]interface{}{"token": "Bearer " + token}})
}
