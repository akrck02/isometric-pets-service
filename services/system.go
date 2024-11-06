package services

import (
	"net/http"

	apierror "github.com/akrck02/isometric-pets-service/error"
	"github.com/akrck02/isometric-pets-service/models"
)

func Health(context *models.ApiContext) (*models.Response, *models.Error) {

	return &models.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func NotImplemented(context *models.ApiContext) (*models.Response, *models.Error) {

	return nil, &models.Error{
		Error:   apierror.NotImplemented,
		Message: "Not implemented",
		Status:  http.StatusNotImplemented,
	}
}

func EmptyCheck(context *models.ApiContext) *models.Error {
	return nil
}
