package routes

import (
	"gogin-fizz/routes/cat"
	"gogin-fizz/routes/dog"

	"github.com/wI2L/fizz"
)

func RegisterRoutes(f *fizz.Fizz) {
	cat.Routes(f)
	dog.Routes(f)
}
