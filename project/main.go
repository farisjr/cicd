package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}
	e := echo.New()
	e.GET("/", hello)
	e.GET("/:name", helloName)
	fmt.Println(port)
	if err := e.Start(port); err != nil {
		fmt.Println(err)
	}
}

func hello(c echo.Context) error {
	return c.HTML(200, "<h1>Hello world</h1><p>Halo dunia, saya bahagia</p>")
}

func helloName(c echo.Context) error {
	name := c.Param("name")
	return c.String(200, fmt.Sprintf("Hello %s", name))
}
