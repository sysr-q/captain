package main

import (
	"github.com/sysr-q/captain/envconfig"
	"github.com/sysr-q/captain/report"
	"github.com/sysr-q/captain/stats"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Server interface {
	Run() error
	Stop() error // Run on SIGINT/SIGTERM. Should do internal cleanup.
}

var (
	config *envconfig.Config = &envconfig.Config{}
	server Server
)

func main() {
	err := envconfig.Process("captain", config)
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(config.Mode) != "stats" &&
		strings.ToLower(config.Mode) != "report" {

		log.Fatal("Given operation mode must be either 'stats' or 'report'.")
	}

	config.Report = strings.ToLower(config.Mode) == "report"

	if config.Report {
		server = report.New(config)
	} else {
		server = stats.New(config)
	}

	signals := make(chan os.Signal)
	// Functionality on Windows not a guarantee.
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Run(); err != nil {
			// Fail early, fail hard.
			log.Fatal(err)
		}
	}()

	<-signals
	if server == nil {
		// This should never happen, but just in case.
		log.Fatal("main.server is nil, can't handle signal.")
		return
	}

	if err := server.Stop(); err != nil {
		log.Fatal(err)
	}

	log.Print("Exiting cleanly, goodbye!")
}
