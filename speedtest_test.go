package speedtest

import (
	"fmt"
	"github.com/rafdekar/speedtest/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkOokla(b *testing.B) {
	for n := 0; n < b.N; n++ {
		speed, err := MeasureSpeed(OOkla)
		if err != nil {
			panic(err)
		}
		fmt.Println(speed)
	}
}

func BenchmarkNetflix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		speed, err := MeasureSpeed(Netflix)
		if err != nil {
			panic(err)
		}
		fmt.Println(speed)
	}
}

func TestMeasureSpeed(t *testing.T) {
	testCases := []struct {
		testName     string
		providerName string
		checkResult  func(t *testing.T, res string, err error)
	}{
		{
			testName:     "Netflix OK",
			providerName: Netflix,
			checkResult: func(t *testing.T, res string, err error) {
				require.NoError(t, err)
				require.True(t, len(res) > 0)
			},
		},
		{
			testName:     "OOkla OK",
			providerName: OOkla,
			checkResult: func(t *testing.T, res string, err error) {
				require.NoError(t, err)
				require.True(t, len(res) > 0)
			},
		},
		{
			testName:     "Invalid Provider",
			providerName: "test",
			checkResult: func(t *testing.T, res string, err error) {
				require.Error(t, err, util.ErrInvalidProvider)
				require.Len(t, res, 0)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			res, err := MeasureSpeed(tc.providerName)
			tc.checkResult(t, res, err)
		})
	}
}
