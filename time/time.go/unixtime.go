package main

import (
    "fmt"
	"time"
)

func main() {
	fmt.Println("🕙\n---")
	
	currentTime := time.Now()

	timeStr := currentTime.UTC().Format(time.RFC3339)
	timeInSec := currentTime.Unix()
	timeInMs := currentTime.UnixMilli()

	fmt.Println(timeStr, "| bash=\"wl-copy", timeStr, "\" terminal=false")
	fmt.Println(timeInSec, "| bash=\"wl-copy", timeInSec, "\" terminal=false")
	fmt.Println(timeInMs, "| bash=\"wl-copy", timeInMs, "\" terminal=false")
	fmt.Println("Epoch Converter | href=https://www.epochconverter.com")
}