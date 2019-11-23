package main

import (
	"fmt"
	"time"
)

func emit(s time.Time, o *options) {
	var prefix string = ""
	var postfix string = "\n"
	if o.fancy {
		prefix = "\r"
		postfix = ""
	}
	fmt.Printf("%s%s%s", prefix, fmtTimer(s, o), postfix)
}

func fmtTimer(s time.Time, o *options) string {
	var d time.Duration
	if o.timer {
		d = o.duration - time.Since(s)
	} else {
		d = time.Since(s)
	}
	return fmtDuration(d)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
