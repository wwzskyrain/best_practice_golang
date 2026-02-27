package main

import (
	"fmt"
	"time"
)

const DateOnlyWithoutHyphen = "20060102"

func validateSyncDate(syncDateStr *string) error {
	synDate, err := time.Parse(DateOnlyWithoutHyphen, *syncDateStr)
	if err != nil {
		return err
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	validDays := 4
	nDaysAgo := today.AddDate(0, 0, -validDays)

	if nDaysAgo.After(synDate) {
		return fmt.Errorf("[invalid syn date] synDate is too early. synDate=%s nDaysAgo=%s",
			synDate.Format(DateOnlyWithoutHyphen),
			nDaysAgo.Format(DateOnlyWithoutHyphen))
	}
	return nil
}

func main() {
	var datePtr *string = nil
	synDate := "20240104"
	datePtr = &synDate
	if err := validateSyncDate(datePtr); err != nil {
		fmt.Println(err)
	}

}
