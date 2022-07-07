package main

import (
	"fmt"
	"time"
)

func main() {
	options := getOptions()
	start := time.Now()
	ticker := time.Tick(options.frequency)
	timeout := time.After(options.duration)
	emit(start, options)
	for {
		select {
		case <-ticker:
			emit(start, options)
		case <-timeout:
			fmt.Println()
			return
		}
	}
}
