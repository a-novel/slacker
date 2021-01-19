package slacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"net/http"
)

func (api *API) Send(m string, b map[string]interface{}, at []map[string]interface{}) *errors.Error {
	fm := map[string]interface{}{"text": m}

	if b != nil {
		fm["blocks"] = b
	}

	if at != nil {
		fm["attachments"] = at
	}

	jsonString, err := json.Marshal(fm)
	if err != nil {
		return errors.New(
			ErrUnexpectedError,
			fmt.Sprintf("unexpected error during post when marshalling data : %s", err.Error()),
		)
	}

	req, err := http.NewRequest("POST", api.WebHook, bytes.NewBuffer(jsonString))
	if err != nil {
		return errors.New(
			ErrUnexpectedError,
			fmt.Sprintf("unexpected error during post when making request : %s", err.Error()),
		)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return errors.New(
			ErrCannotReachSlackServer,
			fmt.Sprintf("cannot reach slack server : %s", err.Error()),
		)
	}

	if resp.StatusCode > 299 {
		return errors.New(
			ErrUnexpectedServerResponse,
			fmt.Sprintf("unexpected response status %v from server", resp.StatusCode),
		)
	}

	return nil
}
