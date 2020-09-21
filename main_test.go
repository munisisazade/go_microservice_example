package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"performance/config"
	"testing"
)

func initConfig() string {
	result := "PORT=:4563\n" +
		"LOG_FILE=exampletest.log\n" +
		"APP_NAME=testservice\n"
	f, err := os.OpenFile(".test",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(result) + "\n");
		err != nil {
		fmt.Println("Cool")
	}
	return "OK"
}

func removeConfig()  {
	e := os.Remove(".test")
	if e != nil {
		log.Fatal(e)
	}
}

func TestConfig(t *testing.T) {
	res := initConfig()
	if res == "OK" {
		e := godotenv.Load(".test") //Load .test file
		if e != nil {
			// if problem problem
			log.Panic(e)
		}
		fmt.Println(config.Name)
		if config.Name != "testservice" {
			log.Panic("Config Name is not valid")
		}
		if config.Port != ":4563" {
			log.Panic("Config Port is not valid")
		}
		if config.Logfile != "exampletest.log" {
			log.Panic("Config Logfile is not valid")
		}
		removeConfig()
	}



}
