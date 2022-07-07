package main

import (
	"flag"
	"time"
)

type options struct {
	fancy     bool
	timer     bool
	duration  time.Duration
	frequency time.Duration
}

func getOptions() *options {
	o := &options{}
	flag.BoolVar(&o.fancy, "fancy", true, "fancy output")
	flag.BoolVar(&o.timer, "timer", false, "start a timer, which counts down")
	flag.DurationVar(&o.duration, "duration", 24*time.Hour, "duration after which the stopwatch stops")
	flag.DurationVar(&o.frequency, "frequency", 1*time.Second, "console update frequency")
	flag.Parse()
	return o
}
