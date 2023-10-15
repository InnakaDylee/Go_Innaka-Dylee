package main

import (
	"Praktikum/controllers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()

	e.POST("/rekomendasi-laptop", controllers.RecLaptopController)

	e.Logger.Fatal(e.Start(":8000"))
}