package main

import (
  "net/http"
  "strconv"

  "github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
func GetUserController(c echo.Context) error {
  // your solution here
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest,map[string]interface{}{
      "massage":"error invalid input",
    })
  }
  var foundUser *User
  for _,user := range users{
    if user.Id == id {
      foundUser = &user
      break
    }
  }
  if foundUser == nil{
    return c.JSON(http.StatusNotFound,map[string]interface{}{
      "massage":"user not found",
    })
  }

  return c.JSON(http.StatusOK,map[string]interface{}{
    "massage":"success get user",
    "user": foundUser,
  })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest,map[string]interface{}{
      "massage":"error invalid input",
    })
  }
  var foundIndex = -1
  for index ,user := range users{
    if user.Id == id {
      foundIndex = index
      break
    }
  }
  if foundIndex == -1{
    return c.JSON(http.StatusNotFound,map[string]interface{}{
      "massage":"user not found",
    })
  }
  if foundIndex != -1{
    users = append(users[:foundIndex],users[foundIndex+1:]... )
  }
  return c.JSON(http.StatusOK,map[string]interface{}{
    "massage":"user delete successfully",
  })

}
// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest,map[string]interface{}{
      "massage":"error invalid input",
    })
  }
  for i ,user := range users{
    if user.Id == id {
      userUpdated := new(User)
      if err := c.Bind(userUpdated); err != nil {
        return err
      }
      users[i].Name = userUpdated.Name
      users[i].Email = userUpdated.Email
      users[i].Password = userUpdated.Password

      return c.JSON(http.StatusOK,map[string]interface{}{
        "massage":"success update user",
        "user": users[i],
      })
    }
  }
  return c.JSON(http.StatusNotFound,map[string]interface{}{
    "massage":"user not found",
  })
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":user,
  })
}
// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.GET("/users/:id", GetUserController)
  e.DELETE("/users/:id", DeleteUserController)
  e.PUT("/users/:id", UpdateUserController)
  e.POST("/users", CreateUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}