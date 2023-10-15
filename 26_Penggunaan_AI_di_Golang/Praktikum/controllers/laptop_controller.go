package controllers

import (
	"Praktikum/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)


func RecLaptopController(c echo.Context) error {
	var input models.Input
	err := c.Bind(&input)
	if err != nil{
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}
	content := fmt.Sprintf("Rekomendasi Laptop budget %s untuk %s merk %s ram %s.", input.Budget, input.Purpose, input.Brand, input.Ram )

	resp, err := RecLaptop(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false,
			Data: "Gagal memberikan rekomendasi",
		})
	}

	data := models.Response{
		Status: true,
		Data: resp,
	}

	return c.JSON(http.StatusOK, data)

}