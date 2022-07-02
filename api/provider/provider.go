package provider

import (
	"github.com/rafdekar/speedtest/api/provider/netflix"
	"github.com/rafdekar/speedtest/api/provider/ookla"
	"github.com/rafdekar/speedtest/util"
)

var (
	OOkla   = &ookla.OOkla{}
	Netflix = &netflix.Netflix{}
)

// Provider interface is defining methods signatures that can be used to calculate user internet speed
type Provider interface {
	Init() error
	GetDownloadSpeed() float64
	GetUploadSpeed() float64
}

// MeasureSpeed uses functionality of provided provider to calculate download and upload rates of user
func MeasureSpeed(provider Provider) (download, upload float64, err error) {
	if provider == nil {
		return 0, 0, util.ErrInvalidProvider
	}

	err = provider.Init()
	if err != nil {
		return 0, 0, err
	}

	return provider.GetDownloadSpeed(), provider.GetUploadSpeed(), nil
}
