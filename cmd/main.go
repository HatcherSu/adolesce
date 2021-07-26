package main

import (
	"cloud_callback/internal/conf"
	"cloud_callback/internal/pkg/log"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read command line
	var filename string
	pflag.StringVar(&filename, "env", ".env", "environment setting file")
	pflag.Parse()
	// read config
	confObj, err := conf.NewConf(filename)
	if err != nil {
		fmt.Printf("NewConf error : %v %v\n", color.RedString("Error:"), err)
	}
	config, err := conf.NewConfig(confObj)
	if err != nil {
		fmt.Printf("NewConfig error :%v %v\n", color.RedString("Error:"), err)
	}
	// initLogger
	logger, err := log.NewLogger(config)
	if err != nil {
		fmt.Printf("NewLogger error :%v %v\n", color.RedString("Error:"), err)
	}
	defer logger.Flush()
	// init application
	app, cleanup, err := initApp(config, logger)
	if err != nil {
		fmt.Printf("initApp error :%v %v\n", color.RedString("Error:"), err)
	}
	app.RegisterStopFunc(cleanup)
	// app.run
	app.Run()
}
