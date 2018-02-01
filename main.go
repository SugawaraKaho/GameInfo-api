package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", save)
	e.GET("/users", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))

}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	return c.String(http.StatusOK, "get dayo")
}

func save(c echo.Context) error {
	// Get name and email
	user := new(User)
	c.Bind(user)

	return c.JSON(http.StatusOK, user)
}
