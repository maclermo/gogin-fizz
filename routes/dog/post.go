package dog

import (
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"

	"gogin-fizz/handlers/dog"
)

func Post(grp *fizz.RouterGroup) {
	grp.POST("", []fizz.OperationOption{
		fizz.Summary("Create a dog"),
		fizz.Response("500", "Server Error", nil, nil, nil),
	}, tonic.Handler(dog.AddDog, 204))
}
