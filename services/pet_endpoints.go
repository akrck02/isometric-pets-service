package services

import "github.com/akrck02/isometric-pets-service/models"

var GetPetEndpoint = models.Endpoint{
	Path:     "/pets/{uuid}",
	Method:   models.GetMethod,
	Listener: GetPet,
}

var GetPetImageEndpoint = models.Endpoint{
	Path:             "/pets/{uuid}/image",
	Method:           models.GetMethod,
	Listener:         GetPetImage,
	ResponseMimeType: models.MimeImagePng,
}
