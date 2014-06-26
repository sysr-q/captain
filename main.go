package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sysr-q/captain/envconfig"
	_ "github.com/sysr-q/captain/report"
	_ "github.com/sysr-q/captain/stats"
	"log"
)

type Config struct {
	// Is this a report server? false == stats, true == report
	Report bool

	// Stats

	Username string // Username to get access to stats server frontend
	Password string // Password to get access to stats server frontend
	Address  string // Address (+ port) to bind to

	// Report

	Stats envconfig.Strings
}

var config Config

func main() {
	err := envconfig.Process("captain", &config)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(config)
}
