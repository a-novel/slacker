package slacker

import "github.com/a-novel/errors"

func (api *API) Verify() *errors.Error {
	if api.WebHook == "" {
		return errors.New(
			ErrMissingWebHook,
			"missing webhook in slack configuration",
		)
	}

	return nil
}
