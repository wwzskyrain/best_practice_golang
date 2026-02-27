package main

import (
	"context"
	"flag"
	"log"
	"time"
)

func buildContext() context.Context {
	var fromTime, toTime string
	flag.StringVar(&fromTime, "from_time", "", "扫描的开始时间")
	flag.StringVar(&toTime, "to_time", "", "扫描的结束时间")
	flag.Parse() // 需要主动parse, 否则无法获取到参数的值.
	ctx := context.Background()
	if len(fromTime) > 0 && len(toTime) > 0 {
		ctx = context.WithValue(ctx, "from_time", fromTime)
		ctx = context.WithValue(ctx, "to_time", toTime)
	}
	return ctx
}

const (
	timeFormat = "2006-01-02 15:04"
)

func main() {
	ctx := buildContext()
	get(ctx)

}

func get(ctx context.Context) {
	now := time.Now()
	fromTime := now.Add(-time.Minute * 15).Format(timeFormat)
	toTime := now.Format(timeFormat)
	if ctxFromTime, ok := ctx.Value("from_time").(string); ok {
		fromTime = ctxFromTime
	}
	if ctxToTime, ok := ctx.Value("to_time").(string); ok {
		toTime = ctxToTime
	}
	log.Printf("fromTime = %s, toTime = %s", fromTime, toTime)
}
