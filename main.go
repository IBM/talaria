package main

import (
	"log/slog"
	"net/http"
	"opentalaria/logger"
	"opentalaria/utils"
	"os"

	// We start a web server only in localdev mode, which should't expose any sensitive information.
	// If we add some web APIs one day, this functionality has to be reviewed.
	_ "expvar"
)

func initLogger() {
	// print the log level before setting the log level handler so we can see what is set in case warn or error are set.
	logLevel := utils.GetLogLevel()
	slog.Info("Setting log level to " + logLevel.String())

	// initialize logger with level handler based on LOG_LEVEL env variable.
	// The default log level is Warn, if no env is set or the value is invalid.
	//
	// JSON Handler might be better suited for a cloud environment. Set it with LOG_FORMAT=json env variable
	var handler slog.Handler
	if os.Getenv("LOG_FORMAT") == "json" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = logger.NewCustomHandler(os.Stdout, nil)
	}

	logger := slog.New(logger.NewLevelHandler(logLevel, handler))

	slog.SetDefault(logger)
}

func main() {
	initLogger()

	if utils.GetProfile() == utils.Localdev {
		slog.Info("starting in local dev mode ...")
		// start a web server if we are in local dev mode
		go http.ListenAndServe(utils.GetEnvVar("DEBUG_SERVER_PORT", ":9090"), nil)
	}

	broker, err := NewBroker()
	if err != nil {
		slog.Error("Error initializing broker", "err", err)
		os.Exit(1)
	}

	if len(broker.Listeners) > 1 {
		slog.Error("OpenTalaria does not support more than one listener for now. See https://github.com/IBM/opentalaria/issues/18")
		os.Exit(1)
	}

	server := NewServer(broker)
	server.Run()
}
