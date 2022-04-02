package dog

import (
	"gogin-fizz/models/dog"

	"github.com/gin-gonic/gin"
)

func GetDog(c *gin.Context, params *dog.DogParams) (dog.Dog, error) {
	return dog.Dog{}, nil
}

func ListDog(c *gin.Context) ([]dog.Dog, error) {
	return []dog.Dog{}, nil
}

func AddDog(c *gin.Context, params *dog.DogParams) (dog.Dog, error) {
	return dog.Dog{}, nil
}
