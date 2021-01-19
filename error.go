package slacker

import (
	"fmt"
	"github.com/a-novel/errors"
	"os"
)

func (api *API) Error(m string) *errors.Error {
	env := os.Getenv("ENV")
	// If no ENV is specified, assume we are in development mode, so we don't want to flood Slack uselessly.
	if env == "" || env == "development" {
		fmt.Println(m)
		return nil
	}

	return api.Send(
		fmt.Sprintf("Unexpected error in %s (%s)", api.Application, env),
		nil,
		[]map[string]interface{}{
			{
				"fallback": fmt.Sprintf("Unexpected error in %s (%s)", api.Application, env),
				"color":    "#FF9300",
				"text":     api.Print(m),
			},
		},
	)
}

func (api *API) Errorf(format string, a ...interface{}) *errors.Error {
	return api.Error(fmt.Sprintf(format, a...))
}
