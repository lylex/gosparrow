package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/tylerb/graceful"

	"lylex.com/gosparrow/app"
	"lylex.com/gosparrow/consts"
	"lylex.com/gosparrow/lib/log"
	"lylex.com/gosparrow/lib/utils"
	"lylex.com/gosparrow/routes"
)

// version represents the application version, and it is valued during build time
var version string

func main() {
	port := os.Getenv(consts.EnvVariablePort)
	if port == "" {
		port = consts.ServiceDefaultPort
	}

	appCtx := &app.Context{
		Version:   version,
		Name:      consts.ServiceName,
		StartedAt: utils.Now(),
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
