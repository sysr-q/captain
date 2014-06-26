package stats

import (
	"errors"
	"github.com/sysr-q/captain/envconfig"
	"log"
	"strings"
)

type Stats struct {
	config *envconfig.Config
}

func (s *Stats) Run() error {
	if strings.TrimSpace(s.config.Address) == "" {
		return errors.New("stats: Must provide a listening address")
	}

	if s.config.Username == "" || s.config.Password == "" {
		log.Printf(
			"You probably forgot to set a username (%q) or password (%q).",
			s.config.Username,
			s.config.Password)
	}

	log.Printf("Listening on %s", s.config.Address)
	return nil
}

func (s *Stats) Stop() error {
	return nil
}

func New(config *envconfig.Config) *Stats {
	return &Stats{
		config: config,
	}
}
