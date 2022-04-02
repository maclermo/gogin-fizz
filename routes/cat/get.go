package cat

import (
	"gogin-fizz/handlers/cat"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

func Get(grp *fizz.RouterGroup) {
	grp.GET("", []fizz.OperationOption{
		fizz.Summary("List all cats"),
		fizz.Response("500", "Server Error", nil, nil, nil),
		fizz.Security(&openapi.SecurityRequirement{
			"basicAuth": []string{},
		}),
	}, tonic.Handler(cat.ListCat, 200))

	grp.GET("/:name", []fizz.OperationOption{
		fizz.Summary("Get one cat"),
		fizz.Response("404", "Not Found", nil, nil, nil),
		fizz.Response("500", "Server Error", nil, nil, nil),
		fizz.Security(&openapi.SecurityRequirement{
			"basicAuth": []string{},
		}),
	}, tonic.Handler(cat.GetCat, 200))
}
