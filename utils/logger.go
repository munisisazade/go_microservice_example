package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"performance/config"
	"performance/models"
	"runtime"
	"strings"
	"time"
)

// File logger 5 status codes
var levels = [5]string{
	"FATAL",
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}

// Create or append file to log object
func write(data models.LogFormat) {
	result, _ := json.Marshal(data)
	f, err := os.OpenFile(config.Logfile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(result) + "\n");
		err != nil {
		fmt.Println("Cool")
	}

}

// all log object standart format
func format(function string, file string, no int, level string, message interface{}) models.LogFormat {
	loc, _ := time.LoadLocation("Asia/Baku")
	now := time.Now().In(loc)
	result := models.LogFormat{
		AppName:   config.Name,
		Function:  function,
		File:      file,
		Line:      no,
		Level:     level,
		Message:   message,
		Timestamp: now,
	}
	return result
}

// logger class use as struct
// usage example
// logger := Logger{}
// logger.Debug("Data fetch !")
type Logger struct {
}

func (i *Logger) Fatal(msg interface{}) {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		file_main := strings.Split(file, "/")
		file = file_main[len(file_main)-1]
		result := format(details.Name(), file, no, levels[0], msg)
		write(result)
	}
}

func (i *Logger) Error(msg interface{}) {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		file_main := strings.Split(file, "/")
		file = file_main[len(file_main)-1]
		result := format(details.Name(), file, no, levels[1], msg)
		write(result)
	}
}

func (i *Logger) Warning(msg interface{}) {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		file_main := strings.Split(file, "/")
		file = file_main[len(file_main)-1]
		result := format(details.Name(), file, no, levels[2], msg)
		write(result)
	}
}

func (i *Logger) Info(msg interface{}) {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		file_main := strings.Split(file, "/")
		file = file_main[len(file_main)-1]
		result := format(details.Name(), file, no, levels[3], msg)
		write(result)
	}
}

func (i *Logger) Debug(msg interface{}) {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		file_main := strings.Split(file, "/")
		file = file_main[len(file_main)-1]
		result := format(details.Name(), file, no, levels[4], msg)
		write(result)
	}
}
