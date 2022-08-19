/*
 * Swagger Petstore
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/p0x0q/go-sample-backend/go"
)

func main() {
	log.Printf("Server started")

	PetsApiService := openapi.NewPetsApiService()
	PetsApiController := openapi.NewPetsApiController(PetsApiService)

	router := openapi.NewRouter(PetsApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
