# Speedtest [![Go Reference](https://pkg.go.dev/badge/github.com/rafdekar/speedtest.svg)][godoc]

### Package to measure internet speed

Provides API to measure internet speed by using either OOkla speedtest.net or Netflix fast.com.

### Installation

Run this command to install the package

    go get github.com/rafdekar/speedtest

### Example use
    
    speed, err := speedtest.MeasureSpeed(speedtest.Netflix)
	if err != nil {
		panic(err)
	}
	fmt.Println(speed)

### Documentation

There is an [online reference for the package][godoc].

[godoc]: https://pkg.go.dev/github.com/rafdekar/speedtest
[golang-install]: http://golang.org/doc/install.html
