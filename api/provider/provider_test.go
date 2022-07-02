package provider

import (
	"github.com/golang/mock/gomock"
	"github.com/rafdekar/speedtest/api/provider/mock"
	"github.com/rafdekar/speedtest/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMeasureSpeed(t *testing.T) {
	testCases := []struct {
		name        string
		passNil     bool
		buildStubs  func(provider *mock.MockProvider)
		checkResult func(dl, ul float64, err error)
	}{
		{
			name:    "OK",
			passNil: false,
			buildStubs: func(provider *mock.MockProvider) {
				provider.EXPECT().Init().Times(1)
				provider.EXPECT().GetDownloadSpeed().Times(1).Return(100.0)
				provider.EXPECT().GetUploadSpeed().Times(1).Return(100.0)
			},
			checkResult: func(dl, ul float64, err error) {
				require.NoError(t, err)
				require.NotZero(t, dl)
				require.NotZero(t, ul)
			},
		},
		{
			name:    "Invalid Provider",
			passNil: true,
			buildStubs: func(provider *mock.MockProvider) {
				provider.EXPECT().Init().Times(0)
				provider.EXPECT().GetDownloadSpeed().Times(0)
				provider.EXPECT().GetUploadSpeed().Times(0)
			},
			checkResult: func(dl, ul float64, err error) {
				require.Error(t, err, util.ErrInvalidProvider)
				require.Zero(t, dl)
				require.Zero(t, ul)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProvider := mock.NewMockProvider(ctrl)

			v.buildStubs(mockProvider)

			var dl, ul float64
			var err error
			if v.passNil {
				dl, ul, err = MeasureSpeed(nil)
			} else {
				dl, ul, err = MeasureSpeed(mockProvider)
			}

			v.checkResult(dl, ul, err)
		})
	}
}
