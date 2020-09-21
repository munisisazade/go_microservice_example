package config

import "os"

var Port = getEnv("PORT", ":8000")
var Logfile = getEnv("LOG_FILE", "micro.log")
var Name = getEnv("APP_NAME", "microservice")

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}