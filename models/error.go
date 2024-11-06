package models

import apierror "github.com/akrck02/isometric-pets-service/error"

type Error struct {
	Status  int               `json:"status,omitempty"`
	Error   apierror.ApiError `json:"error,omitempty"`
	Message string            `json:"message,omitempty"`
}
