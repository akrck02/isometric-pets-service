package common

import (
	"net/http"
	"os"

	"github.com/akrck02/isometric-pets-service/configuration"
	"github.com/akrck02/isometric-pets-service/middleware"
	"github.com/akrck02/isometric-pets-service/models"
	"github.com/akrck02/isometric-pets-service/services"
	"github.com/withmandala/go-log"
)

const API_PATH = "/"

var logger = log.New(os.Stderr)

// ApiMiddlewares is a list of middleware functions that will be applied to all API requests
// this list can be modified to add or remove middlewares
// the order of the middlewares is important, it will be applied in the order they are listed
var ApiMiddlewares = []middleware.Middleware{
	middleware.Trazability,
}

func Start(configuration configuration.APIConfiguration, endpoints []models.Endpoint) {

	// set debug or release mode
	if configuration.IsDevelopment() {
		logger.WithDebug()
	}

	// show log app title and start router
	logger.Info(configuration.ApiName)

	// Add API path to endpoints
	newEndpoints := []models.Endpoint{}
	for _, endpoint := range endpoints {
		endpoint.Path = API_PATH + configuration.ApiName + "/" + configuration.Version + endpoint.Path
		newEndpoints = append(newEndpoints, endpoint)
	}

	// Register endpoints
	registerEndpoints(newEndpoints)

	// Start listening HTTP requests
	logger.Infof("API started on http://%s:%s%s%s/", configuration.Ip, configuration.Port, API_PATH, configuration.ApiName)
	logger.Info("")
	state := http.ListenAndServe(configuration.Ip+":"+configuration.Port, nil)
	logger.Error(state.Error())

}

func registerEndpoints(endpoints []models.Endpoint) {

	for _, endpoint := range endpoints {

		switch endpoint.Method {
		case models.GetMethod:
			endpoint.Path = "GET " + endpoint.Path
		case models.PostMethod:
			endpoint.Path = "POST " + endpoint.Path
		case models.PutMethod:
			endpoint.Path = "PUT " + endpoint.Path
		case models.DeleteMethod:
			endpoint.Path = "DELETE " + endpoint.Path
		case models.PatchMethod:
			endpoint.Path = "PATCH " + endpoint.Path
		}

		logger.Infof("Endpoint %s registered.", endpoint.Path)

		// set defaults
		setEndpointDefaults(&endpoint)

		// register endpoint
		http.HandleFunc(endpoint.Path, func(writer http.ResponseWriter, request *http.Request) {

			// log the request
			logger.Info("")
			logger.Infof("%s", endpoint.Path)

			// enable CORS
			writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))
			writer.Header().Set("Access-Control-Allow-Methods", os.Getenv("CORS_METHODS"))
			writer.Header().Set("Access-Control-Allow-Headers", os.Getenv("CORS_HEADERS"))
			writer.Header().Set("Access-Control-Max-Age", os.Getenv("CORS_MAX_AGE"))

			// create basic api context
			context := &models.ApiContext{
				Trazability: models.Trazability{
					Endpoint: endpoint,
				},
			}

			// Get request data
			err := middleware.Request(request, context)
			if nil != err {
				middleware.SendResponse(writer, err.Status, err, models.MimeApplicationJson)
				return
			}

			// Apply middleware to the request
			err = applyMiddleware(context)
			if nil != err {
				middleware.SendResponse(writer, err.Status, err, models.MimeApplicationJson)
				return
			}

			// Execute the endpoint and send the response
			middleware.Response(context, writer)
		})

	}
}

func setEndpointDefaults(endpoint *models.Endpoint) {

	if nil == endpoint.Checks {
		endpoint.Checks = services.EmptyCheck
	}

	if nil == endpoint.Listener {
		endpoint.Listener = services.NotImplemented
	}

	if endpoint.RequestMimeType == "" {
		endpoint.RequestMimeType = models.MimeApplicationJson
	}

	if endpoint.ResponseMimeType == "" {
		endpoint.ResponseMimeType = models.MimeApplicationJson
	}

}

func applyMiddleware(context *models.ApiContext) *models.Error {

	for _, middleware := range ApiMiddlewares {
		err := middleware(context)
		if nil != err {
			return err
		}
	}

	return nil

}
