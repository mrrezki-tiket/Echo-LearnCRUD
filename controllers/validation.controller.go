package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

type Customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"required,gte=17,lte=35"`
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email := "Rezki"
	err := v.Var(email, "required,email")

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama:   "Rezki",
		Email:  "rezki@mail.com",
		Alamat: "Kediri",
		Umur:   17,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}
