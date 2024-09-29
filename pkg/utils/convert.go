package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-typesense-search/internal/pkg/response"
	"log"
)

func BodyToDto[T any](c *gin.Context) *T {
	var payload *T
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		response.ValidateErrorResponse(c, err)
		c.Abort()
		return nil
	}
	return payload
}

func DtoToModel[MT any, T any](dto T) (*MT, *response.ServerCode) {
	var model MT

	bytes, _ := json.Marshal(dto)
	err := json.Unmarshal(bytes, &model)
	if err != nil {
		log.Println("convert to dto failed", err.Error())
		return nil, response.ReturnCode(response.ErrCodeParamInvalid)
	}

	return &model, nil
}

func ModelToDto[T any, MT any](model MT) *T {
	var payload T

	bytes, _ := json.Marshal(model)
	err := json.Unmarshal(bytes, &payload)
	if err != nil {
		log.Printf("convert to dto failed %v", err.Error())
		return nil
	}
	return &payload
}
