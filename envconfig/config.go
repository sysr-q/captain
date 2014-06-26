// This file contains a special struct type that's referenced by various modules
// in Captain. So we just keep it in captain/envconfig for ease of use.

package envconfig

type Config struct {
	// Is this a report server? false == stats, true == report
	Report bool

	// Stats

	Username string // Username to get access to stats server frontend
	Password string // Password to get access to stats server frontend
	Address  string // Address (+ port) to bind to

	// Report

	Stats Strings
}
