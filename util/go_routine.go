package util

import (
	"context"
	"log"
	"runtime/debug"
)

func SafeGoRoutine(ctx context.Context, identifier string, f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[%s] panic: %+v", identifier, err)
				log.Printf("[%s] Stack trace: %+v", identifier, string(debug.Stack()))
			}
		}()
		log.Printf("[%s] start...", identifier)
		f()
		log.Printf("[%s] finished.", identifier)
	}()
}
