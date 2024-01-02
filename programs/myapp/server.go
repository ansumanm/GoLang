package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	t:= &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = t
	e.GET("/hello", Hello)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Routing
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

// Path Parameters
func getUser(c echo.Context) error {
	// User ID from path 'users/:id'
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// Query Parameters
// /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam(("member"))
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

// curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "name" + name + ", email:" + email)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "Templates")
}