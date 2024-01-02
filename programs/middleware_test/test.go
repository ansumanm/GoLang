package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Middleware function
func myMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Executing middleware...")
		return next(c)
	}
}

// Handler function
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	// Use the middleware
	e.Use(myMiddleware)

	// Define a route and attach the handler
	e.GET("/", helloHandler)

	// Start the server
	e.Start(":8080")
}
