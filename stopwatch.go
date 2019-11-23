package main

import (
	"flag"
	"time"
)

func main() {
	options := getOptions()
	if options.help {
		flag.PrintDefaults()
		return
	}
	start := time.Now()
	ticker := time.Tick(options.frequency)
	timeout := time.After(options.duration)
	emit(start, options)
	for {
		select {
		case <-ticker:
			emit(start, options)
		case <-timeout:
			return
		}
	}
}
