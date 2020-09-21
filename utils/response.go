package utils

import (
	"github.com/satori/go.uuid"
	"net/http"
	"performance/config"
	"performance/models"
	"time"
)

func NotFoundResponse() models.Response {
	uuid_string := uuid.Must(uuid.NewV4())
	loc, _ := time.LoadLocation("Asia/Baku")
	now := time.Now().In(loc)
	maintime := now.Format("2006-01-02 15:04:05")
	response := models.Response{
		Timestamp:   maintime,
		Status:      404,
		Description: http.StatusText(404),
		Data:        nil,
		Transaction: uuid_string,
		Exception:   models.Exception{
			Code: "PAGE_NOT_FOUND",
			Message: "Səhifə tapılmadı.",
		},
		AppName:     config.Name,
	}
	return response
}

func SuccessResponse(data interface{}) models.Response {
	uuid_string := uuid.Must(uuid.NewV4())
	loc, _ := time.LoadLocation("Asia/Baku")
	now := time.Now().In(loc)
	maintime := now.Format("2006-01-02 15:04:05")
	response := models.Response{
		Timestamp:   maintime,
		Status:      200,
		Description: http.StatusText(200),
		Data:        data,
		Transaction: uuid_string,
		Exception:   nil,
		AppName:     config.Name,
	}
	return response
}