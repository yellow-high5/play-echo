package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func syslogFormat() string {
	var format string

	format += "${time_rfc3339}  "
	format += "${status} "
	format += "${method} "
	format += "${host}${uri} "
	format += "[user_agent:${user_agent}] "
	format += "(remote_ip:${remote_ip}, "
	format += "req_size:${bytes_in}B, "
	format += "res_size:${bytes_out}B, "
	format += "request_time:${latency_human})\n"
	return format
}

func DebugLogFormat() string {
	var format string

	format += "${time_rfc3339}  "
	format += "${level}  "
	format += "(${short_file}:${line})"
	return format
}

func SetLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: syslogFormat(),
		Output: os.Stdout,
	})
}

func SetBodyDump(l echo.Logger) echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if len(reqBody) != 0 {
			l.Infof("Request Body: %v\n", string(reqBody))
		}
		if len(resBody) != 0 {
			l.Infof("Response Body: %v\n", string(resBody))
		}
	})
}
