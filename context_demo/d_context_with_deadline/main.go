package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	// Create a context with a deadline of 1 second.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Wait for the operation to complete or the context to expire.
	select {
	case <-time.After(2 * time.Second):
		// Operation completed within the deadline
		fmt.Println("Operation completed")
	case <-ctx.Done():
		// Context deadline exceeded
		err := ctx.Err()
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Context deadline exceeded")
		} else if errors.Is(err, context.Canceled) {
			fmt.Println("Context canceled")
		} else {
			fmt.Println("Unknown context error:", err)
		}
	}
}
