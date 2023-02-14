package router

import (
	"net/http"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

type routerConfig struct {
	Middlewares []func(http.Handler) http.Handler
}

func NewRouterConfig(ctx *core.MifyServiceContext) *routerConfig {
	return &routerConfig{
		Middlewares: []func(http.Handler) http.Handler{
			// Add your middlewares here
		},
	}
}
