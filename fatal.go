package slacker

import (
	"fmt"
	"log"
	"os"
)

func (api *API) Fatal(m string) {
	fm := api.Print(m)

	env := os.Getenv("ENV")
	// If no ENV is specified, assume we are in development mode, so we don't want to flood Slack uselessly.
	if env != "" && env != "development" {
		_ = api.Send(
			fmt.Sprintf("Unexpected error in %s (%s)", api.Application, env),
			nil,
			[]map[string]interface{}{
				{
					"fallback": fmt.Sprintf("Fatal error in %s (%s)", api.Application, env),
					"color":    "#FF3232",
					"text":     fm,
				},
			},
		)
	}

	log.Fatalf(fm)
}

func (api *API) Fatalf(format string, a ...interface{}) {
	api.Fatal(fmt.Sprintf(format, a...))
}
