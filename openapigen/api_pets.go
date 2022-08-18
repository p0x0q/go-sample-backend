/*
 * Swagger Petstore
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// PetsApiController binds http requests to an api service and writes the service results to the http response
type PetsApiController struct {
	service PetsApiServicer
	errorHandler ErrorHandler
}

// PetsApiOption for how the controller is set up.
type PetsApiOption func(*PetsApiController)

// WithPetsApiErrorHandler inject ErrorHandler into controller
func WithPetsApiErrorHandler(h ErrorHandler) PetsApiOption {
	return func(c *PetsApiController) {
		c.errorHandler = h
	}
}

// NewPetsApiController creates a default api controller
func NewPetsApiController(s PetsApiServicer, opts ...PetsApiOption) Router {
	controller := &PetsApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PetsApiController
func (c *PetsApiController) Routes() Routes {
	return Routes{ 
		{
			"ListPets",
			strings.ToUpper("Get"),
			"/v1/pets",
			c.ListPets,
		},
	}
}

// ListPets - List all pets
func (c *PetsApiController) ListPets(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ListPets(r.Context(), limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}