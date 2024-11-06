package services

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"

	apierror "github.com/akrck02/isometric-pets-service/error"
	"github.com/akrck02/isometric-pets-service/models"
)

func GetPet(context *models.ApiContext) (*models.Response, *models.Error) {
	var uuid string = context.Request.Params["uuid"]
	if uuid == "" {
		return nil, &models.Error{
			Error:   apierror.BadRequest,
			Message: "uuid is required",
			Status:  http.StatusBadRequest,
		}
	}

	return &models.Response{
		Code: http.StatusOK,
		Response: &models.Pet{
			Uuid: uuid,
			Name: "Teko",
		},
	}, nil

}

func GetPetImage(context *models.ApiContext) (*models.Response, *models.Error) {
	var uuid string = context.Request.Params["uuid"]
	if uuid == "" {
		return nil, &models.Error{
			Error:   apierror.BadRequest,
			Message: "uuid is required",
			Status:  http.StatusBadRequest,
		}
	}

	var images = readImage(uuid)
	return &models.Response{
		Code:     http.StatusOK,
		Response: images,
		Length:   len(images),
	}, nil

}

func readImage(uuid string) []byte {

	f, err := os.Open("resources/" + uuid + ".png")
	if err != nil {
		return nil
	}

	defer f.Close()
	image, _, err := image.Decode(f)

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, image); err != nil {
		log.Println("unable to encode image.")
	}

	if err != nil {
		return nil
	}

	return buffer.Bytes()
}
