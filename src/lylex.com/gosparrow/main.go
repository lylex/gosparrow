package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/tylerb/graceful"

	"lylex.com/audit/app"
	"lylex.com/audit/consts"
	"lylex.com/audit/lib/discovery"
	"lylex.com/audit/lib/log"
	"lylex.com/audit/lib/utils"
	"lylex.com/audit/routes"
	"lylex.com/audit/workers"
)

// version represents the application version, and it is valued during build time
var version string

func main() {
	port := os.Getenv(consts.EnvVariablePort)
	if port == "" {
		port = consts.ServiceDefaultPort
	}

	discovery := discovery.New()
	localDiscovery, globalDiscovery, err := discovery.RetrieveDiscoveryConfigs()
	if err != nil {
		log.App(consts.LogLevelError,
			fmt.Sprintf(consts.ErrFailInitDiscovery, err.Error()))
	}

	appCtx := &app.Context{
		Version:   version,
		Name:      consts.ServiceName,
		StartedAt: utils.Now(),
	}

	if err = log.InitLog(appCtx.LocalDiscovery.LogLevel); err != nil {
		log.App(consts.LogLevelError,
			fmt.Sprintf("Failed to set log level: %s",
				err.Error(),
			),
		)
	}

	app.SetContext(appCtx.Assimilate())

	server := negroni.New()
	server.UseHandler(routes.NewRouter())

	log.App(consts.LogLevelInfo,
		fmt.Sprintf("%s service run on port %s...", consts.ServiceName, port))
	if err := graceful.RunWithErr(":"+port, 10*time.Second, server); err != nil {
		log.App(consts.LogLevelError,
			fmt.Sprintf("Error occurs: %s, will exit",
				err.Error(),
			),
		)
		os.Exit(1)
	}

	log.App(consts.LogLevelInfo,
		fmt.Sprintf("%s service exits", consts.ServiceName))
}
