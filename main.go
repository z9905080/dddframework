package main

import (
	ApiServer "dddframework/api_server"
	EnvVar "dddframework/environment/env_variable"
	"flag"
	Logging "github.com/z9905080/gloger"
	"os"
)

// init init for program
func init() {

	Logging.SetCurrentLevel(Logging.DEBUG)
	Logging.SetLogMode(Logging.Stdout)
	Logging.Force("current environment:", os.Getenv("ENV"))

	// Set Log Mode & Level
	Logging.SetCurrentLevel(Logging.Level(EnvVar.Setting.LogLevel))
	Logging.SetLogMode(Logging.OutputMode(EnvVar.Setting.LogMode))

	//if EnvVar.Setting.Debug {
	//	Logging.Debug("Service run on DEBUG mode")
	//}
}

const (
	CMD_API_SERVER = "api"
)

func main() {

	var command string
	flag.StringVar(&command, "cmd", "default", "api: run an instance with API Service")
	flag.Parse()

	switch command {
	case CMD_API_SERVER:
		{
			Logging.Force("Start API Server... ")
			ApiServer.Run()
		}
	default:
		{
			Logging.Error("not match cmd, will exit...")
		}

	}

}
