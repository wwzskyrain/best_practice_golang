package main

import (
	"fmt"
	"time"
)

func main() {
	// parsedT, err := time.Parse("20230114", "20240108")
	//
	dateStr := "20240106"
	parsedT, err := time.Parse("20060102", dateStr)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("parsed parsedT:", parsedT)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println("today:", today)
	threeDayBefore := today.AddDate(0, 0, -3)

	fmt.Println("threeDayBefore:", threeDayBefore)

	// 减10天后，就是2023年了，哈哈哈
	tenDayBefore := today.AddDate(0, 0, -10)
	fmt.Println("tenDayBefore:", tenDayBefore)

	fmt.Println(today.After(parsedT))

	for i := 0; i < 10; i++ {
		updatedDay := today.AddDate(0, 0, -i)
		fmt.Printf("updatedDay:%s , 早于: %s\n", updatedDay.Format(time.DateOnly), today.Format(time.DateOnly))
	}
}
