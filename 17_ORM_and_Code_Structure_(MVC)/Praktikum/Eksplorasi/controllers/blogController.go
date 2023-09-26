package controllers

import (
	"Prioritas_1_2/configs"
	"Prioritas_1_2/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	result := configs.DB.Preload("User").Find(&blogs)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":blogs,
	})
}

func GetBlogController(c echo.Context) error {
	var blogs []models.Blog

	id,_ := strconv.Atoi(c.Param("id"))

	blog := configs.DB.Preload("User").Find(&blogs, id)
	if err := blog.Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if blog != nil {
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"users":   blog.Value,
	  })
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "blog not found",
	})
}

func CreateBlogController(c echo.Context) error {
	var blog models.Blog
	c.Bind(&blog)

	if err := configs.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	var blogs models.Blog

	id,_ := strconv.Atoi(c.Param("id"))

	blog := configs.DB.Preload("User").Find(&blogs, id)
	if err := blog.Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if blog != nil {
	  configs.DB.Delete(&blogs, id)
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "delete blog success",
		"user": blogs,
	  })
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "blog not found",
	})
}

  // update user by id
  func UpdateBlogController(c echo.Context) error {
	// your solution here
	var blog models.Blog
	var edit models.Blog
	c.Bind(&edit)
  
	id,_ := strconv.Atoi(c.Param("id"))
	
	if err := configs.DB.Find(&blog).Error ;err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"massage": "fail get blog",
	  })
	}
	configs.DB.Preload("User").Find(&blog, id)
	blog.UsersId = edit.UsersId
	blog.Title = edit.Title
	blog.Content = edit.Content
  
	configs.DB.Save(&blog)
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "massage": "update blog success",
	  "user" : blog,
	})
}