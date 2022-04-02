package dog

import (
	"github.com/wI2L/fizz"
)

type RouteConfig struct {
	Path        string
	Name        string
	Description string
}

func Routes(f *fizz.Fizz) {

	dogConfig := RouteConfig{
		Path:        "/dog",
		Name:        "dog",
		Description: "Dog operations",
	}
	dogGroup := f.Group(dogConfig.Path, dogConfig.Name, dogConfig.Description)

	Get(dogGroup)
	Post(dogGroup)
}
