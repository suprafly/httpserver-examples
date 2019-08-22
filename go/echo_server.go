package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  e := echo.New()
  e.Pre(middleware.RemoveTrailingSlash())
  e.GET("/users", getUsers)
	e.GET("/users/:user_id", getUser)
  e.Logger.Fatal(e.Start(":8080"))
}

func getUsers(c echo.Context) error {
  data := []int{}
  return c.JSON(http.StatusOK, data)
}

func getUser(c echo.Context) error {
	user_id := c.Param("user_id")
	return c.String(http.StatusOK, user_id)
}
