package slacker

import (
	"fmt"
	"github.com/a-novel/errors"
	"os"
)

func (api *API) Log(m, n, c string) *errors.Error {
	env := os.Getenv("ENV")
	// If no ENV is specified, assume we are in development mode, so we don't want to flood Slack uselessly.
	if env == "" || env == "development" {
		fmt.Println(m)
		return nil
	}

	return api.Send(
		fmt.Sprintf("%s (%s)", n, api.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback": fmt.Sprintf("%s (%s)", n, api.Application),
				"color":    c,
				"text":     api.Print(m),
			},
		},
	)
}
