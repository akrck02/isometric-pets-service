package middleware

import (
	"github.com/akrck02/isometric-pets-service/models"
	"github.com/akrck02/isometric-pets-service/utils"
)

func Trazability(context *models.ApiContext) *models.Error {

	time := utils.GetCurrentMillis()

	context.Trazability = models.Trazability{
		Endpoint:  context.Trazability.Endpoint,
		Timestamp: &time,
	}

	return nil
}
