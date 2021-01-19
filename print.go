package slacker

import (
	"fmt"
	"runtime/debug"
	"time"
)

func (api *API) Print(m string) string {
	return fmt.Sprintf(
		"*Message*\n%s\n\n*Stack*\n```%s```\n\n*Time*\n%s",
		m,
		string(debug.Stack()),
		time.Now().Format("2006-01-02 3:4:5"),
	)
}
