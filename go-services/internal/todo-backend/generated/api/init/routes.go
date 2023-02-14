// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify via OpenAPI Generator

// vim: set ft=go:
package openapi_init

import (
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/api"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
	"github.com/go-chi/chi/v5"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/handlers/todos"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/handlers/todos/id"
)

func Routes(ctx *core.MifyServiceContext, routerConfig openapi.RouterConfig, reqExtraFactory core.RequestExtraFactory) chi.Router {

	TodosApiService := api_todos.NewTodosApiService()
	TodosApiController := openapi.NewTodosApiController(ctx, TodosApiService)

	TodosIdApiService := api_todos_id.NewTodosIdApiService()
	TodosIdApiController := openapi.NewTodosIdApiController(ctx, TodosIdApiService)

	router := openapi.NewRouter(ctx, routerConfig, reqExtraFactory, TodosApiController, TodosIdApiController)
	return router
}