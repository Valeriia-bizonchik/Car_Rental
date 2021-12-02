package main

import (
	"fmt"
	"github.com/Valeriia-bizonchik/CarRental/internal/api/rest/v1"
	"os"
	"os/signal"
	"syscall"

	"github.com/Valeriia-bizonchik/CarRental/config"
	"github.com/Valeriia-bizonchik/CarRental/internal/storage/mysql"
	"github.com/Valeriia-bizonchik/CarRental/logger"
)

func main() {
	cfg, err := config.InitEnvConfig()
	if err != nil {
		fmt.Println(`failed to parse config, system exit`)
		os.Exit(1)
	}

	zLog := logger.InitZapFileConsole(cfg.DebugMode, cfg.LogFile)
	defer zLog.Sync()

	storage, err := mysql.NewCarRentalStorage(cfg.MySqlDNS)
	if err != nil {
		zLog.Sugar().Error(err)
	}

	api := v1.NewAPI(storage, zLog.Sugar())
	api.InitRoutes()

	errs := make(chan error, 1)

	go func() {
		errs <- api.Run(cfg.ServiceHost, cfg.ServicePort)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	zLog.Sugar().Error("terminated: ", <-errs)
}
