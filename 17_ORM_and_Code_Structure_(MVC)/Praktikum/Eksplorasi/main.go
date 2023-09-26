package main

import (
  "Prioritas_1_2/controllers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)


func main() {
  // create a new echo instance
  e := echo.New()
  // Route / to handler function
  e.GET("/users", controllers.GetUsersController)
  e.GET("/users/:id", controllers.GetUserController)
  e.POST("/users", controllers.CreateUserController)
  e.DELETE("/users/:id", controllers.DeleteUserController)
  e.PUT("/users/:id", controllers.UpdateUserController)

  e.GET("/blogs", controllers.GetBlogsController)
  e.POST("/blogs", controllers.CreateBlogController)
  e.GET("/blogs/:id", controllers.GetBlogController)
  e.DELETE("/blogs/:id", controllers.DeleteBlogController)
  e.PUT("/blogs/:id", controllers.UpdateBlogController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}