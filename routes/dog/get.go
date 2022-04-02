package dog

import (
	"gogin-fizz/handlers/dog"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

func Get(grp *fizz.RouterGroup) {
	grp.GET("", []fizz.OperationOption{
		fizz.Summary("List all dogs"),
		fizz.Response("500", "Server Error", nil, nil, nil),
	}, tonic.Handler(dog.ListDog, 200))

	grp.GET("/:name", []fizz.OperationOption{
		fizz.Summary("Get one dog"),
		fizz.Response("404", "Not Found", nil, nil, nil),
		fizz.Response("500", "Server Error", nil, nil, nil),
	}, tonic.Handler(dog.GetDog, 200))
}
