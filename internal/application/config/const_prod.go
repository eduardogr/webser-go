//go:build prod
// +build prod

package config

const (
	// get config from external provider
	STRATEGY = STRATEGY_EXTERNAL

	API_PORT = 8080
)
