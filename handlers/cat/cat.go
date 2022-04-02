package cat

import (
	"gogin-fizz/models/cat"

	"github.com/gin-gonic/gin"
)

func GetCat(c *gin.Context, params *cat.CatParams) (cat.Cat, error) {
	return cat.Cat{}, nil
}

func ListCat(c *gin.Context) ([]cat.Cat, error) {
	return []cat.Cat{}, nil
}

func AddCat(c *gin.Context, params *cat.CatParams) (cat.Cat, error) {
	return cat.Cat{}, nil
}
