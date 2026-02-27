package time_util

import (
	"fmt"
	"time"
)

func PrintDuration(ms1, ms2 int64) {
	// 转成 time.Time
	t1 := time.UnixMilli(ms1)
	t2 := time.UnixMilli(ms2)

	// 打印原始毫秒值
	fmt.Printf("t1 毫秒值: %d\n", ms1)
	fmt.Printf("t2 毫秒值: %d\n", ms2)

	// 打印对应的年月日时分秒
	fmt.Printf("t1 时间: %s\n", t1.Format("2006年01月02日 15时04分05秒"))
	fmt.Printf("t2 时间: %s\n", t2.Format("2006年01月02日 15时04分05秒"))

	// 计算 Duration
	var duration time.Duration
	if ms1 > ms2 {
		duration = time.Duration(ms1-ms2) * time.Millisecond
	} else {
		duration = time.Duration(ms2-ms1) * time.Millisecond
	}

	// 打印 Duration 的年月日时分秒形式
	// 注意：Duration 并不是一个绝对时间，需要手动拆分
	totalSeconds := int64(duration.Seconds())
	seconds := totalSeconds % 60
	minutes := (totalSeconds / 60) % 60
	hours := (totalSeconds / 3600) % 24
	days := totalSeconds / (3600 * 24)

	fmt.Printf("Duration: %d天 %02d时 %02d分 %02d秒\n", days, hours, minutes, seconds)
}
