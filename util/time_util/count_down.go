package time_util

import (
	"log"
	"time"
)

// CountDownToZero 倒计时
func CountDownToZero(name string, n int) {
	for n > 0 {
		time.Sleep(time.Second)
		log.Printf("[%s] 倒计时 【%d】", name, n)
		n--
	}
}
