# Slacker

A Slack api for Go, to redirect critical server errors in real time in a Slack channel.

```bash
go get github.com/a-novel/slacker
```

# Prerequisite

You need a [Slack application](https://api.slack.com/apps) that supports [incoming webhooks](https://api.slack.com/messaging/webhooks#:~:text=Incoming%20Webhooks%20are%20a%20simple,make%20the%20messages%20stand%20out.).

# Configuration

Create a JSON file with 2 keys : your Slack channel webhook and an optional application name.

```json
{
  "webhook":     "https://hooks.slack.com/services/XXXXX/XXXXX/XXXXXXX",
  "application": "My Super App"
}
```

# Init

Once your configuration file is ready, you can run the init function in your main go file:

```go
package myPackage

import (
	"github.com/a-novel/slacker"
	"log"
)

func init() {
	logger, err := slacker.InitFromFile("/path/to/configuration/file.json")
	if err != nil {
		log.Fatal(err.Error())
	}
}
```

# Handle errors

Once your client is ready, you can use your logger object to send preformatted messages to your Slack channel.

```go
func main() {
    // Do stuff..
	// An error occurs......
	
	logger.Errorf("an error occurred: %s", err.Error())
	// You can also use logger.Fatalf() to shut down your server because of a critical error.
}
```

# Usage with Gin

You can use Slacker to automatically catch gin aborted requests:

```go
package myPackage

import (
	"github.com/a-novel/slacker"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	logger.UseGinFormatter(r)
}
```

# License
2021, A-Novel [MIT License](https://github.com/a-novel/slacker/blob/master/LICENSE).