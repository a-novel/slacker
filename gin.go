package slacker

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (api *API) GinFormatter(param gin.LogFormatterParams) string {
	err := api.Error(
		fmt.Sprintf(
			"*%s* %s [%v]\n\n```%s```\n\n",
			param.Method,
			param.Path,
			param.StatusCode,
			param.ErrorMessage,
		),
	)

	if err != nil {
		fmt.Println(err.Error())
	}

	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func (api *API) UseGinFormatter(router *gin.Engine) {
	router.Use(gin.LoggerWithFormatter(api.GinFormatter))
}
