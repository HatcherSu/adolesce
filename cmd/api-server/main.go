package main

import (
	"adolesce/internal/conf"
	"adolesce/pkg/log"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"math/rand"
	"runtime"
	"time"
)

var (
	filename = pflag.String("env", ".env", "environment setting file")
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read command line
	pflag.Parse()
	// read config
	confObj, err := conf.NewConf(*filename)
	if err != nil {
		fmt.Printf("NewConf error : %v %v\n", color.RedString("Error:"), err)
		return
	}
	config, err := conf.NewConfig(confObj)
	if err != nil {
		fmt.Printf("NewConfig error :%v %v\n", color.RedString("Error:"), err)
		return
	}
	// initLogger
	logger, err := log.NewLogger(config)
	if err != nil {
		fmt.Printf("NewLogger error :%v %v\n", color.RedString("Error:"), err)
		return
	}
	defer logger.Flush()
	// init application
	app, cleanup, err := initApp(config, logger)
	if err != nil {
		fmt.Printf("initApp error :%v %v\n", color.RedString("Error:"), err)
		return
	}
	app.RegisterStopFunc(cleanup)
	// app.run
	app.Run()
}