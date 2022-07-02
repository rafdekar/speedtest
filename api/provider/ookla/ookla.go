package ookla

import (
	"github.com/rafdekar/speedtest/util"
	"github.com/showwin/speedtest-go/speedtest"
)

// OOkla is struct for www.speedtest.net website
type OOkla struct {
	userData *speedtest.User
	server   *speedtest.Server
}

// Init method is downloading user info, fetching the closest servers and doing download, and upload tests to the selected servers
func (o *OOkla) Init() error {
	var err error

	o.userData, err = speedtest.FetchUserInfo()
	if err != nil {
		return util.ErrCouldNotFetchUserInfo
	}

	servers, err := speedtest.FetchServers(o.userData)
	if err != nil {
		return util.ErrCouldNotFetchServers
	}
	if len(servers) <= 0 {
		return util.ErrServersNotFound
	}

	o.server = findClosestServer(servers)

	err = o.server.DownloadTest(true)
	if err != nil {
		return util.ErrDownloadTestFailure
	}

	err = o.server.UploadTest(false)
	if err != nil {
		return util.ErrUploadTestFailure
	}

	return nil
}

// GetDownloadSpeed returns download speed of a user
func (o *OOkla) GetDownloadSpeed() float64 {
	return o.server.DLSpeed
}

// GetUploadSpeed returns upload speed of a user
func (o *OOkla) GetUploadSpeed() float64 {
	return o.server.ULSpeed
}

// findClosestServer function is selecting the closes server from range of servers
func findClosestServer(servers speedtest.Servers) (server *speedtest.Server) {
	distance := 100000.0
	for _, v := range servers {
		if v.Distance < distance {
			distance = v.Distance
			server = v
		}
	}
	return
}
