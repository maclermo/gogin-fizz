package cat

import (
	"gogin-fizz/middlewares"

	"github.com/wI2L/fizz"
)

type RouteConfig struct {
	Path        string
	Name        string
	Description string
}

func Routes(f *fizz.Fizz) {
	catConfig := RouteConfig{
		Path:        "/cat",
		Name:        "cat",
		Description: "Cat operations",
	}
	catGroup := f.Group(catConfig.Path, catConfig.Name, catConfig.Description)
	catGroup.Use(middlewares.BasicAuthMiddleware())

	Get(catGroup)
	Post(catGroup)
}
