// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Pokemon Studio API
 *
 * API for the Pokemon Studio
 *
 * API version: 0.0.1
 */

package psapigen

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// PokemonAPIController binds http requests to an api service and writes the service results to the http response
type PokemonAPIController struct {
	service      PokemonAPIServicer
	errorHandler ErrorHandler
}

// PokemonAPIOption for how the controller is set up.
type PokemonAPIOption func(*PokemonAPIController)

// WithPokemonAPIErrorHandler inject ErrorHandler into controller
func WithPokemonAPIErrorHandler(h ErrorHandler) PokemonAPIOption {
	return func(c *PokemonAPIController) {
		c.errorHandler = h
	}
}

// NewPokemonAPIController creates a default api controller
func NewPokemonAPIController(s PokemonAPIServicer, opts ...PokemonAPIOption) *PokemonAPIController {
	controller := &PokemonAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PokemonAPIController
func (c *PokemonAPIController) Routes() Routes {
	return Routes{
		"GetPokemon": Route{
			strings.ToUpper("Get"),
			"/api/pokemon",
			c.GetPokemon,
		},
		"GetPokemonDetails": Route{
			strings.ToUpper("Get"),
			"/api/pokemon/{symbol}",
			c.GetPokemonDetails,
		},
		"GetPokemonForm": Route{
			strings.ToUpper("Get"),
			"/api/pokemon/{symbol}/{form}",
			c.GetPokemonForm,
		},
	}
}

// GetPokemon - Get a page of pokemon
func (c *PokemonAPIController) GetPokemon(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var pageParam int32
	if query.Has("page") {
		param, err := parseNumericParameter[int32](
			query.Get("page"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](0),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "page", Err: err}, nil)
			return
		}

		pageParam = param
	} else {
		var param int32 = 0
		pageParam = param
	}
	var sizeParam int32
	if query.Has("size") {
		param, err := parseNumericParameter[int32](
			query.Get("size"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](50),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "size", Err: err}, nil)
			return
		}

		sizeParam = param
	} else {
		var param int32 = 20
		sizeParam = param
	}
	acceptLanguageParam := r.Header.Get("Accept-Language")
	result, err := c.service.GetPokemon(r.Context(), pageParam, sizeParam, acceptLanguageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPokemonDetails - Get a pokemon details
func (c *PokemonAPIController) GetPokemonDetails(w http.ResponseWriter, r *http.Request) {
	symbolParam := chi.URLParam(r, "symbol")
	if symbolParam == "" {
		c.errorHandler(w, r, &RequiredError{"symbol"}, nil)
		return
	}
	acceptLanguageParam := r.Header.Get("Accept-Language")
	result, err := c.service.GetPokemonDetails(r.Context(), symbolParam, acceptLanguageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPokemonForm - Get a pokemon form details
func (c *PokemonAPIController) GetPokemonForm(w http.ResponseWriter, r *http.Request) {
	symbolParam := chi.URLParam(r, "symbol")
	if symbolParam == "" {
		c.errorHandler(w, r, &RequiredError{"symbol"}, nil)
		return
	}
	formParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "form"),
		WithDefaultOrParse[int32](0, parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Param: "form", Err: err}, nil)
		return
	}
	acceptLanguageParam := r.Header.Get("Accept-Language")
	result, err := c.service.GetPokemonForm(r.Context(), symbolParam, formParam, acceptLanguageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
