package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/tylerb/graceful"

	"gosparrow/app"
	"gosparrow/consts"
	"gosparrow/lib/log"
	"gosparrow/lib/utils"
	"gosparrow/routes"
)

// version represents the application version, and it is valued during build time
var version string

func main() {
	// Allow gorutines in all the cpus, go by default makes it runs in one
	runtime.GOMAXPROCS(runtime.NumCPU())

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
