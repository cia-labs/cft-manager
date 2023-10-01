package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cia-labs/cft-manager/internal"
)

var (
	Version   = "DEV"
	BuildTime = "unknown"
	BuildType = ""
)

func main() {
	configFile := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("Error: config file path is required")
		os.Exit(1)
	}

	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		fmt.Printf("Error: config file %s not found\n", *configFile)
		os.Exit(1)
	}

	logger := internal.GetLogger()
	logger.Info("Starting CFT Manager")
}
