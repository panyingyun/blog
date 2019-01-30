//go:generate go-bindata -prefix ./public/ -o=assets/public/public_gen.go -pkg=public public/...
//go:generate go-bindata -prefix ./templates/ -o=assets/templates/templates_gen.go -pkg=templates templates/...

package main

import (
	"blog/applog"
	"blog/config"
	"blog/httpserver"

	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func run(c *cli.Context) error {
	conf, err := config.ReadConfig(c.String("conf"))
	if err != nil {
		log.Error("read from conf fail!", c.String("conf"))
		return err
	}
	fmt.Println("conf =  ", conf)

	fmt.Println("runtime.GOOS = ", runtime.GOOS)

	var logger *applog.AutoDailyLoger
	if runtime.GOOS == "windows" {
		logger = applog.NewAutoDailyLoger(conf.LogDirWin, conf.LogPrefix)
	} else {
		logger = applog.NewAutoDailyLoger(conf.LogDirLinux, conf.LogPrefix)
	}
	logger.Start()
	defer logger.Stop()

	//start http server
	go func() {
		if runtime.GOOS == "windows" {
			httpserver.StartHttpServer(conf.HttpServerWin)
		} else {
			httpserver.StartHttpServer(conf.HttpServerLinux)
		}
	}()

	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Infof("signal received signal %v", <-sigChan)
	log.Warn("shutting down server")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "michaelapp"
	app.Usage = "michaelapp blog"
	app.Copyright = "panyingyun@gmail.com"
	app.Version = "0.0.1"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "conf,c",
			Usage:  "Set conf path here",
			Value:  "appserver_prod.conf",
			EnvVar: "APP_CONF",
		},
	}
	app.Run(os.Args)
}
