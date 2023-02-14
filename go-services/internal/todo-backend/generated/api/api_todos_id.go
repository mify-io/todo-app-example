// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify via OpenAPI Generator

// vim: set ft=go:
package openapi

import (
	"net/http"
	"strings"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/api/public"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

// TodosIdApiController binds http requests to an api service and writes the service results to the http response
type TodosIdApiController struct {
	ctx          *core.MifyServiceContext
	service      TodosIdApiServicer
	errorHandler ErrorHandler
}

// TodosIdApiOption for how the controller is set up.
type TodosIdApiOption func(*TodosIdApiController)

// WithTodosIdApiErrorHandler inject ErrorHandler into controller
func WithTodosIdApiErrorHandler(h ErrorHandler) TodosIdApiOption {
	return func(c *TodosIdApiController) {
		c.errorHandler = h
	}
}

// NewTodosIdApiController creates a default api controller
func NewTodosIdApiController(ctx *core.MifyServiceContext, s TodosIdApiServicer, opts ...TodosIdApiOption) Router {
	controller := &TodosIdApiController{
		ctx:          ctx,
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the TodosIdApiController
func (c *TodosIdApiController) Routes() Routes {
	return Routes{
		{
			"TodosIdDelete",
			strings.ToUpper("Delete"),
			"/todos/{id}",
			c.TodosIdDelete,
		},
		{
			"TodosIdGet",
			strings.ToUpper("Get"),
			"/todos/{id}",
			c.TodosIdGet,
		},
		{
			"TodosIdPut",
			strings.ToUpper("Put"),
			"/todos/{id}",
			c.TodosIdPut,
		},
	}
}

// TodosIdDelete - Delete todo note by id
func (c *TodosIdApiController) TodosIdDelete(w http.ResponseWriter, r *http.Request) {
	var handlerErr error
	var requestBody []byte
	idParamRaw, err := parseInt64Parameter(getURLParam(r, "id"), true)
	if err != nil {
		handlerErr = err
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	idParam := idParamRaw

	reqCtx := openapi_public.GetMifyRequestContext(r)

	var result ServiceResponse
	defer LogHandler(reqCtx, requestBody, &result, &handlerErr)()

	var herr error
	result, herr = c.service.TodosIdDelete(reqCtx, idParam)
	// If an error occurred, encode the error with the status code
	if herr != nil {
		handlerErr = herr
		c.errorHandler(w, r, herr, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// TodosIdGet - Get todo note by id
func (c *TodosIdApiController) TodosIdGet(w http.ResponseWriter, r *http.Request) {
	var handlerErr error
	var requestBody []byte
	idParamRaw, err := parseInt64Parameter(getURLParam(r, "id"), true)
	if err != nil {
		handlerErr = err
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	idParam := idParamRaw

	reqCtx := openapi_public.GetMifyRequestContext(r)

	var result ServiceResponse
	defer LogHandler(reqCtx, requestBody, &result, &handlerErr)()

	var herr error
	result, herr = c.service.TodosIdGet(reqCtx, idParam)
	// If an error occurred, encode the error with the status code
	if herr != nil {
		handlerErr = herr
		c.errorHandler(w, r, herr, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// TodosIdPut - Update todo note
func (c *TodosIdApiController) TodosIdPut(w http.ResponseWriter, r *http.Request) {
	var handlerErr error
	var requestBody []byte
	idParamRaw, err := parseInt64Parameter(getURLParam(r, "id"), true)
	if err != nil {
		handlerErr = err
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	idParam := idParamRaw

	todoNoteUpdateRequestParamRaw := TodoNoteUpdateRequest{}
	var rerr error
	requestBody, rerr = readBody(r.Body)
	if rerr != nil {
		handlerErr = rerr
		c.errorHandler(w, r, &ParsingError{Err: rerr}, nil)
		return
	}
	d := getBodyDecoder(requestBody)
	d.DisallowUnknownFields()
	if err := d.Decode(&todoNoteUpdateRequestParamRaw); err != nil {
		handlerErr = err
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTodoNoteUpdateRequestRequired(todoNoteUpdateRequestParamRaw); err != nil {
		handlerErr = err
		c.errorHandler(w, r, err, nil)
		return
	}

	todoNoteUpdateRequestParam := todoNoteUpdateRequestParamRaw

	reqCtx := openapi_public.GetMifyRequestContext(r)

	var result ServiceResponse
	defer LogHandler(reqCtx, requestBody, &result, &handlerErr)()

	var herr error
	result, herr = c.service.TodosIdPut(reqCtx, idParam, todoNoteUpdateRequestParam)
	// If an error occurred, encode the error with the status code
	if herr != nil {
		handlerErr = herr
		c.errorHandler(w, r, herr, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}