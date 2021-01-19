package slacker

import (
	"github.com/a-novel/anogo"
	"github.com/a-novel/errors"
)



func InitFromFile(configPath string) (*API, *errors.Error) {
	var output API

	if err := anogo.ReadConfig(configPath, &output); err != nil {
		return nil, err
	}

	return Init(&output)
}

func Init(config *API) (*API, *errors.Error) {
	return config, config.Verify()
}
