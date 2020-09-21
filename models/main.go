package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type LogFormat struct {
	AppName   string      `json:"app_name"`
	Function  string      `json:"function"`
	File      string      `json:"file"`
	Line      int         `json:"line"`
	Level     string      `json:"level"`
	Message   interface{} `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
}

type Response struct {
	AppName     string      `json:"app_name"`
	Data        interface{} `json:"data"`
	Description string      `json:"description"`
	Exception   interface{} `json:"exception"`
	Status      int         `json:"status"`
	Timestamp   string      `json:"timestamp"`
	Transaction uuid.UUID   `json:"transaction"`
}

type Exception struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
