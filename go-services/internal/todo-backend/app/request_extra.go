package app

import (
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

type RequestExtra struct {
	// Append your dependencies here
}

func NewRequestExtra(ctx *core.MifyServiceContext) (*RequestExtra, error) {
	// Here you can do your custom service initialization, prepare dependencies
	extra := &RequestExtra{
		// Here you can initialize your dependencies
	}
	return extra, nil
}
