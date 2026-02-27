package main

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKey("AuthToken"), "X@cfs645JHSDcdae")

	go processSensitiveData(ctx)

	// give some time for other goroutine to conclude
	time.Sleep(2 * time.Second)
}

func processSensitiveData(ctx context.Context) {
	authToken, ok := ctx.Value(contextKey("AuthToken")).(string)
	if !ok {
		// print appropriate message
		fmt.Println("Access Denied - missing AuthToken!")
		return
	}

	fmt.Printf("Processing data using: %s", authToken)
}
