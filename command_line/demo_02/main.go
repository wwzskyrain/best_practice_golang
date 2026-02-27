package main

import (
	"context"
	"flag"
	"log"
	"time"
)

//	cd /opt/tiger/pangle_risk_i18n/monitor_pangle_penalty && ./bootstrap.sh --from_time "2020-05-01 00:00" --to_time "2024-05-10 00:00"
//
// cd /opt/tiger/pangle_risk_i18n/monitor_pangle_penalty && ./bootstrap.sh --from_time "2020-05-01 00:00" --to_time "2024-05-10 00:00"
// --from_time "2020-05-01 00:00" --to_time "2024-05-10 00:00"
func buildContext() context.Context {
	timeFormat := "2006-01-02 15:04"
	now := time.Now()
	defaultFromTime := now.Add(-time.Minute * 15).Format(timeFormat)
	defaultToTime := now.Format(timeFormat)

	var fromTime, toTime string
	flag.StringVar(&fromTime, "from_time", defaultFromTime, "扫描的开始时间")
	flag.StringVar(&toTime, "to_time", defaultToTime, "扫描的结束时间")
	flag.Parse() // 需要主动parse, 否则无法获取到参数的值.
	ctx := context.Background()

	ctx = context.WithValue(ctx, "from_time", fromTime)
	ctx = context.WithValue(ctx, "to_time", toTime)

	return ctx
}

func main() {
	// ctx := buildContext()
	// log.Printf("from_time = %s, to_time = %s", ctx.Value("from_time"), ctx.Value("to_time"))

	// model := parseRunModel()
	// log.Printf("mode = %s\n", model)
	// model, fromTime, toTime := parseTimeWindow()
	// log.Printf("model=%s, fromTime=%s, toTime=%s \n", model, fromTime, toTime)
	runModel, fromTime, toTime := parseParamFromCommandLine()
	if runModel == "test" {
		log.Printf("runModel=%s", runModel)
	}
	log.Printf("fromTime=%s, toTime=%s", fromTime, toTime)

}
func parseParamFromCommandLine() (runModel, fromTime, toTime string) {
	flag.StringVar(&runModel, "run_model", "", "运行模式")
	flag.StringVar(&fromTime, "from_time", "", "扫描的开始时间")
	flag.StringVar(&toTime, "to_time", "", "扫描的结束时间")
	flag.Parse() // 需要主动parse, 否则无法获取到参数的值.
	return
}
