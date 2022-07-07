# Stopwatch

## What is Stopwatch?

Stopwatch is a simple _command line toy_ to display a terminal based stopwatch or timer.

Here is a quick demo of it in action:

![Stopwatch demo](/etc/demo.gif)

## Installation

### Binaries

1) download the binary for the appropriate platform, and architecture
2) add the binary to your system path (optional)

### Compiling from source

If there are no binaries for your platform, or you want to compile from source you can follow these steps:

1) make sure that you have __go version 1.13__ or newer
2) fetch and build the source:

```
go install github.com/garricasaurus/stopwatch@latest
```

3) the binary will be placed in your `$GOPATH/bin/`

## Usage

### Help

Starting the application with the `-help` switch prints out all possible flags and their default values:

```
stopwatch -help

  -duration duration
        duration after which the stopwatch stops (default 24h0m0s)
  -fancy
        fancy output (default true)
  -frequency duration
        console update frequency (default 1s)
  -timer
        start a timer, which counts down
```

### Stopwatch mode

The default mode is stopwatch mode. The program will simply count the seconds and print the elapsed time to the console.

```
stopwatch
```

The stopwatch will stop when you hit `CTRL+C` or after 24 hours (this value can be changed with the `-duration` flag). Duration flags can be provided in a user friendly format as seen below:

```
stopwatch -duration=12m30s
```

### Timer mode

Timer mode can be enabled with the `-timer` flag. In this mode the application will act as a countdown timer starting from the value provided with the `-duration` flag (or 24 hours by default).

```
stopwatch -timer -duration=15s
```

### Display frequency

By default the application updates the elapsed time on the console every second. This can be changed with the `-frequency` flag:

```
stopwatch -frequency=5s
```

### Non-Fancy mode

When printing to the console Stopwatch will try to update the elapsed time on the same line without spamming your console ("_fancy mode_"). If this feature does not work properly on your terminal, or you want to switch it off for some other reason, you can use the `-fancy=false` switch. You can also turn it off in order to redirect the output of Stopwatch.

```
stopwatch -duration=10s -frequency=2s -fancy=false > file.txt

cat file.txt

00:00:00
00:00:02
00:00:04
00:00:06
00:00:08
00:00:10
```