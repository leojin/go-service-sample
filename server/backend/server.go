package server

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	server "github.com/leojin/go-service-bootstrap-server-echo"
)

func NewServer() *server.HTTPServerEcho {

	s := &server.HTTPServerEcho{
		Name: "backend",
		Port: 8012,
	}
	s.SetLogger(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	})

	go func() {
		time.AfterFunc(2*time.Second, func() {
			_ = s.Echo.Shutdown(context.Background())
		})
	}()
	return s
}
