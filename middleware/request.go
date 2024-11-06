package middleware

import (
	"net/http"
	"regexp"

	apierror "github.com/akrck02/isometric-pets-service/error"
	"github.com/akrck02/isometric-pets-service/models"
)

const USER_AGENT_HEADER = "User-Agent"
const AUTHORITATION_HEADER = "Authorization"

func Request(r *http.Request, context *models.ApiContext) *models.Error {

	parserError := r.ParseForm()

	if parserError != nil {
		return &models.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.InvalidRequest,
			Message: parserError.Error(),
		}
	}

	context.Request = models.Request{
		Authorization: r.Header.Get(AUTHORITATION_HEADER),
		IP:            r.Host,
		UserAgent:     r.Header.Get(USER_AGENT_HEADER),
		Headers:       map[string]string{},
		Body:          r.Body,
		Params:        map[string]string{},
	}

	// If files are present, add them to the context
	if r.MultipartForm != nil {
		context.Request.Files = r.MultipartForm.File
	}

	// Add headers to the context
	for key, value := range r.Header {
		for _, v := range value {
			context.Request.Headers[key] = v
		}
	}

	// Add params to the context
	for key, value := range r.Form {
		for _, v := range value {
			context.Request.Params[key] = v
		}
	}

	// Get possible url path parameters
	pathParams := getPathParamNames(context.Trazability.Endpoint.Path)
	for _, param := range pathParams {
		context.Request.Params[param] = r.PathValue(param)
	}

	return nil
}

func getPathParamNames(path string) []string {
	params := []string{}

	//regex to find path parameters
	regex, err := regexp.Compile("{(.*?)}")

	if err != nil {
		return params
	}

	params = regex.FindAllString(path, -1)

	if params == nil {
		params = []string{}
	}

	for i, param := range params {
		params[i] = param[1 : len(param)-1]
	}

	return params
}