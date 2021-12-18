package main

import (
	"flag"
	"log"
	"mock.services/client/controller"
	"mock.services/client/service"
	"mock.services/client/transport"
	"mock.services/common/config"
	"mock.services/common/logger"
	"mock.services/common/server"
	"mock.services/common/utility"
	"os"

	"go.uber.org/zap"
)

var (
	configFile string
)

func parseParameters() {
	flag.StringVar(&configFile, "config", "", "the config file used to launch the service")
	flag.Parse()

	if len(configFile) == 0 {
		log.Fatalf("missing config file")
	}
}

func main() {
	// parseParameters()

	zapLogger, err := logger.SetupLogger(config.Logger{
		Level:     "INFO",
		Filepath:  "./logs/clientProxy.log",
		MaxSize:   128,
		MaxBackup: 4,
		MaxAge:    30,
		Compress:  false,
	})
	if err != nil {
		log.Fatal(err)
	}

	svr := server.NewServer(zapLogger)
	if err := svr.SetupServer(config.Server{
		Address: ":10000",
	}); err != nil {
		zapLogger.Info("setup server failed", zap.Error(err))
		os.Exit(1)
	}

	proxyService := service.NewProxyService(zapLogger)
	proxyController := controller.NewProxyController(zapLogger, proxyService)
	transport.SetupProxyRouters(svr, proxyController)

	zapLogger.Info("start server")

	g := utility.MakeGroup()
	g.Add(svr.RunServer, svr.StopServer)

	if err := g.Run(); err != nil {
		zapLogger.Info("run failed", zap.Error(err))
		os.Exit(1)
	}
}
