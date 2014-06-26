package report

import (
	"errors"
	"github.com/sysr-q/captain/envconfig"
	"log"
	"strings"
)

type Report struct {
	config *envconfig.Config
}

func (r *Report) Run() error {
	if len(r.config.Stats) < 1 {
		return errors.New("report: Must report to one or more stats server")
	}

	plural := ""
	if len(r.config.Stats) > 1 {
		plural = "s"
	}
	log.Printf(
		"Reporting to %d server%s: %s",
		len(r.config.Stats),
		plural,
		strings.Join([]string(r.config.Stats), ", "))
	return nil
}

func (r *Report) Stop() error {
	return nil
}

func New(config *envconfig.Config) *Report {
	return &Report{
		config: config,
	}
}
