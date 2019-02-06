package config

import (
	"github.com/gorilla/mux"
	"github.com/plicca/summerfish-swagger"
	"log"
	"path/filepath"
)

func GenerateSwaggerDocsAndEndpoints(router *mux.Router, endpoint string) (err error) {
	config := summerfish.Config{
		Schemes:                []string{"http", "https"},
		SwaggerFileRoute:       summerfish.SwaggerFileRoute,
		SwaggerFilePath:        summerfish.SwaggerFileRoute,
		SwaggerFileHeaderRoute: summerfish.SwaggerFileRoute,
		SwaggerUIRoute:         summerfish.SwaggerUIRoute,
		BaseRoute:              "/",
	}

	config.SwaggerFilePath, err = filepath.Abs("res/swagger.json")
	if err != nil {
		return
	}

	routerInformation, err := summerfish.GetInfoFromRouter(router)
	if err != nil {
		return
	}

	scheme := summerfish.SchemeHolder{Schemes: config.Schemes, Host: endpoint, BasePath: config.BaseRoute}
	err = scheme.GenerateSwaggerFile(routerInformation, config.SwaggerFilePath)
	if err != nil {
		return
	}

	log.Println("Swagger documentation generated")
	return summerfish.AddSwaggerUIEndpoints(router, config)
}

