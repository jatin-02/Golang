package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Println(currTime.Format("01-02-2006 15:04:05 Monday"))
}