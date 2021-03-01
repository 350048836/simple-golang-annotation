package amyxframework

import (
	"./controller"
	"./environment"
	"./factory/controllerfactory"
	"./factory/propertiesfactory"
	"./httpserver"
	"./utils/logger"
	"./utils/stringutils"
)

const (
	defaultConfigFile string = "application.properties"
)

type StartParams struct {
	PropertiesFile string
	Application    interface{}
	APIs           []interface{}
}

func Start(opts StartParams) {
	logger.Info("AmyxFramework Start")

	if stringutils.IsEmpty(opts.PropertiesFile) {
		opts.PropertiesFile = defaultConfigFile
	}

	propertiesfactory.BuildProperties(opts.PropertiesFile)
	//内部api
	opts.APIs = append(opts.APIs, controller.HealthCheck)
	controllerfactory.BuildController(opts.APIs)

	httpserver.Start(environment.GetProperty("server.port"), controllerfactory.GetRouter())
}
