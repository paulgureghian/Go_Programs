package main

import (
	"context"
	"time"
)

func main() {
	// initialize some resources
	db, err := database.Initialize()
	server, err := http.Initialize()

	// wait for termination signal and register database & http server clean-up operations
	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return db.Shutdown()
		},
		"http-server": func(ctx context.Context) error {
			return server.Shutdown()
		},
		// Add other cleanup operations here
	})

	<-wait
}
