package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	fmt.Println(time)

	date := time.Format("02-01-2006")
	fmt.Println(date)
	date = time.Format("02-01-2006 Monday")
	fmt.Println(date)
	date = time.Format("02-01-2006 15:04:05 Monday")
	fmt.Println(date)

	createtime := time.Date(202)
}
