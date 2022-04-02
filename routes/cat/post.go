package cat

import (
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"

	"gogin-fizz/handlers/cat"
)

func Post(grp *fizz.RouterGroup) {
	grp.POST("", []fizz.OperationOption{
		fizz.Summary("Create a cat"),
		fizz.Response("500", "Server Error", nil, nil, nil),
		fizz.Security(&openapi.SecurityRequirement{
			"basicAuth": []string{},
		}),
	}, tonic.Handler(cat.AddCat, 204))
}
