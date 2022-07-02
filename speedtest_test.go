package main

import (
	"fmt"
	"github.com/rafdekar/speedtest/api/provider"
	"github.com/rafdekar/speedtest/util"
	"testing"
)

func BenchmarkOokla(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dl, ul, err := provider.MeasureSpeed(provider.OOkla)
		if err != nil {
			panic(err)
		}
		fmt.Println(util.FormatMeasures(dl, ul))
	}
}

func BenchmarkNetflix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dl, ul, err := provider.MeasureSpeed(provider.Netflix)
		if err != nil {
			panic(err)
		}
		fmt.Println(util.FormatMeasures(dl, ul))
	}
}
