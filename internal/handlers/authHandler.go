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
)

func RegisterUser(w http.ResponseWriter, r *http.Request){

	var validate *validator.Validate

	var req types.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors){
			fmt.Fprintf(w, "Validation failed for field '%s' : '%s\n'", err.Field(), err.Tag())
		}
		return 
	}

	var existingUser models.User

	if err := database.DB.Where("email=?", req.Email).First(&existingUser).Error; err != nil{
		http.Error(w, "User already exists", http.StatusConflict)
		return 
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		return 
	}


}