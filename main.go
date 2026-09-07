package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	year := now.Year()
	if now.Month() == time.December && now.Day() == 31 {
		fmt.Println("С Новым годом!")
		return
	}
	if now.Month() == time.December || now.Month() > time.January {
		year++
	}
	
	newYear := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	
	days := int(newYear.Sub(now).Hours() / 24)
	
	fmt.Printf("До Нового года осталось %d дней!\n", days)
}