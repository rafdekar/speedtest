package util

import "fmt"

// FormatMeasures formats download and upload data to a nice string
func FormatMeasures(download, upload float64) string {
	return fmt.Sprintf("Download speed: %f\nUpload speed: %f\n", download, upload)
}
