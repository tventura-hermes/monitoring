package main

import (
	"context"
	errors_helpers "demo/helpers/errors"
	"demo/router"
	"demo/trigger"
)

func main() {
	var r router.Routes
	r.CreateRoutes()
	r.StartServer()

	ctx := context.Background()

	receiver, err := trigger.NewTrigger(ctx)
	if err != nil {
		errors_helpers.ReportError(ctx, err)
		panic("trigger setup failed")
	}

	receiver.GetMessage()
}
