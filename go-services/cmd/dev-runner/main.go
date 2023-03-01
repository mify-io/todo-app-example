// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify

package main



import (
	"context"
	todo_backend "example.com/namespace/todo-app/go-services/internal/todo-backend/generated/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	app_todo_backend := todo_backend.NewMifyServiceApp(ctx)
	go func() {
		app_todo_backend.Run()
		cancel()
	}()

	<-ctx.Done()
}


