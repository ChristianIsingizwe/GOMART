package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChristianIsingizwe/GOMART/internal/database"
	"github.com/ChristianIsingizwe/GOMART/internal/helpers"
	"github.com/ChristianIsingizwe/GOMART/internal/models"
	"github.com/ChristianIsingizwe/GOMART/internal/types"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	validate := validator.New()

	validate.RegisterValidation("strongpassword", helpers.StrongPassword)

	var req types.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Fprintf(w, "Validation failed for field '%s' : '%s\n'", err.Field(), err.Tag())
		}
		return
	}

	var existingUser models.User

	if err := database.DB.Where("email=?", req.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

		} else {
			http.Error(w, "Failed to check if the user exists", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		return
	}

	if req.Role == "" {
		req.Role = "customer"
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      req.Role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to register the user", http.StatusInternalServerError)
		return
	}

	accessToken, err := helpers.GenerateAccessToken(fmt.Sprint(user.ID), fmt.Sprint(user.Role), int(user.TokenVersion))
	if err != nil {
		http.Error(w, "Failed to generate the access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := helpers.GenerateRefreshToken(fmt.Sprint(user.ID), fmt.Sprint(user.Role), int(user.TokenVersion))
	if err != nil {
		http.Error(w, "Failed to generate the refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"firstname":     user.FirstName,
		"lastname":      user.LastName,
		"email":         user.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	validate := validator.New()

	var req types.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payloads", http.StatusBadRequest)
		return
	}

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Fprintf(w, "Validation failed for field %s: %s\n", err.Field(), err.Tag())
		}
		return
	}

	var user models.User

	if err := database.DB.Where("email=?", req.Email).First(&user).Error; err == nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := helpers.CheckPassword(req.Password, user.Password); err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	accessToken, err := helpers.GenerateAccessToken(fmt.Sprint(user.ID), fmt.Sprint(user.Role), int(user.TokenVersion))
	if err != nil {
		http.Error(w, "Failed to create the token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := helpers.GenerateRefreshToken(fmt.Sprint(user.ID), fmt.Sprint(user.Role), int(user.TokenVersion))
	if err != nil {
		http.Error(w, "Failed to create the refresh token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"email":         user.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
