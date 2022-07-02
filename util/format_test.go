package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFormatMeasures(t *testing.T) {
	dl, ul := 200.0, 100.0

	formattedSpeeds := FormatMeasures(dl, ul)

	require.Equal(t, "Download speed: 200.00  Upload speed: 100.00", formattedSpeeds)
}
