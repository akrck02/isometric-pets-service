package main

import (
	"github.com/akrck02/isometric-pets-service/common"
	"github.com/akrck02/isometric-pets-service/configuration"
	"github.com/akrck02/isometric-pets-service/models"
	"github.com/akrck02/isometric-pets-service/services"
)

const ENVIROMENT_FILE = ".env"

// main function
func main() {

	config := configuration.LoadConfiguration(ENVIROMENT_FILE)

	common.Start(
		config,
		[]models.Endpoint{
			services.GetPetEndpoint,
			services.GetPetImageEndpoint,
		},
	)

}
