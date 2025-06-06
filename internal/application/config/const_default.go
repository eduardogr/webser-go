//go:build !dev && !prod
// +build !dev,!prod

package config

const (
	STRATEGY = STRATEGY_FROM_CODE

	DB_USER     = "root"
	DB_PASSWORD = "testpass"
	DB_HOST     = "mysql" // mysql container
	DB_DATABASE = "basedb"
	DB_PORT     = 3306 // container expose 3308 but 3306 should be use within the docker's network

	API_PORT = 8080
)
