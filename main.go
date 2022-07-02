package main

import (
	"fmt"
	"github.com/rafdekar/speedtest/api/provider"
	"github.com/rafdekar/speedtest/util"
)

func main() {
	download, upload, err := provider.MeasureSpeed(provider.Netflix)
	if err != nil {
		return
	}

	fmt.Println(util.FormatMeasures(download, upload))
}
