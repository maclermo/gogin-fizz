package main

import (
	"net/http"

	"gogin-fizz/middlewares"
	"gogin-fizz/routes"

	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

func main() {
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(middlewares.CORSMiddleware())

	f := fizz.NewFromEngine(engine)

	infos := &openapi.Info{
		Title:       "Sample boilerplate",
		Description: "This is my sample fizz boilerplate app.",
		Version:     "1.0.0",
	}

	securityBasicAuthScheme := &openapi.SecuritySchemeOrRef{
		SecurityScheme: &openapi.SecurityScheme{
			Type:   "http",
			Scheme: "basic",
		},
	}

	f.Generator().API().Components.SecuritySchemes = map[string]*openapi.SecuritySchemeOrRef{
		"basicAuth": securityBasicAuthScheme,
	}

	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))

	routes.RegisterRoutes(f)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: f,
	}

	srv.ListenAndServe()
}
