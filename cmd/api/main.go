package main

import (
	"github.com/ChristianIsingizwe/GOMART/internal/helpers"
	"github.com/go-playground/validator/v10"
)

func main() {
	var validate *validator.Validate = validator.New()
	validate.RegisterValidation("strongpassword", helpers.StrongPassword)

}
