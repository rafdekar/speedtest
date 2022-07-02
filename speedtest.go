package speedtest

import (
	"github.com/rafdekar/speedtest/api/provider"
	"github.com/rafdekar/speedtest/util"
)

const (
	OOkla   = "ookla"
	Netflix = "netflix"
)

// MeasureSpeed measures internet speed of user using specified provider. Provider can be either Netflix or OOkla.
func MeasureSpeed(providerName string) (string, error) {
	switch providerName {
	case OOkla:
		dl, ul, err := provider.MeasureSpeed(provider.OOkla)
		if err != nil {
			return "", err
		}
		return util.FormatMeasures(dl, ul), nil
	case Netflix:
		dl, ul, err := provider.MeasureSpeed(provider.Netflix)
		if err != nil {
			return "", err
		}
		return util.FormatMeasures(dl, ul), nil
	default:
		return "", util.ErrInvalidProvider
	}
}
