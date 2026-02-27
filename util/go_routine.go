package util

import (
	"context"
	"runtime/debug"

	"code.byted.org/gopkg/logs"
)

func SafeGoRoutine(ctx context.Context, identifier string, f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.CtxError(ctx, "[%s] panic: %+v", identifier, err)
				logs.CtxError(ctx, "[%s] Stack trace: %+v", identifier, string(debug.Stack()))
			}
		}()
		logs.CtxInfo(ctx, "[%s] start...", identifier)
		f()
		logs.CtxInfo(ctx, "[%s] finished.", identifier)
	}()
}
