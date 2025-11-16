package logger

import (
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is the global logger instance
var Logger *logrus.Logger

func InitLogger(debug bool) {
	// Ensure logs folder exists
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	// Lumberjack log rotation
	logFile := &lumberjack.Logger{
		Filename:   "logs/api.log",
		MaxSize:    5, // MB
		MaxBackups: 7,
		MaxAge:     30, // days
		Compress:   true,
	}

	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	Logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
	Logger.SetReportCaller(true)

	if debug {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
}

// Echo middleware for request logging
func EchoLoggerMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			latency := time.Since(start)

			Logger.WithFields(logrus.Fields{
				"method":  c.Request().Method,
				"uri":     c.Request().RequestURI,
				"status":  c.Response().Status,
				"latency": latency,
			}).Info("HTTP request")
			return err
		}
	}
}
