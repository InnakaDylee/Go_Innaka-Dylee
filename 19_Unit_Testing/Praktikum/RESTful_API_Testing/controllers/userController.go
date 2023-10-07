package controllers

import (
	"Praktikum/configs"
	"Praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)
// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User
  
	if err := configs.DB.Find(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
  }
  // get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var user models.User
	var users []models.User
  
	id, _ := strconv.Atoi(c.Param("id"))
	if err := configs.DB.First(&users).Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userFound := configs.DB.First(&user, id)
	userValue := configs.DB.First(&users, id)
	if userFound.Error == nil {
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"users":   userValue.Value,
	})
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}
  // create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := configs.DB.Save(&user).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new user",
	  "user":    user,
	})
}
  // delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var user models.User
	var users []models.User
	id, _ := strconv.Atoi(c.Param("id"))
  
	if err := configs.DB.Find(&users).Error ;err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"massage": "fail get user",
	  })
	}
	userFound := configs.DB.First(&user, id)
	userValue := configs.DB.First(&users, id)
	if userFound.Error == nil{
	  configs.DB.Delete(&users, id)
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "delete user success",
		"user": userValue,
	  })
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
	  "massage": "user not found",
	})
}
  
  
  // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var edit models.User
	c.Bind(&edit)
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
  
	if err := configs.DB.Find(&user).Error ;err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"massage": "fail get user",
	  })
	}
	userFound := configs.DB.First(&user, id)
	if userFound.Error != nil{
	  return c.JSON(http.StatusNotFound, map[string]interface{}{
		"massage": "user not found",
	  })
	}
	user.Name = edit.Name
	user.Email = edit.Email
	user.Password = edit.Password
  
	configs.DB.Save(&user)
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "massage": "update user success",
	  "user" : user,
	})
}