package middleware

import "github.com/akrck02/isometric-pets-service/models"

type Middleware func(context *models.ApiContext) *models.Error
