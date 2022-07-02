package util

import "errors"

// Possible errors that can be returned by this library
var (
	ErrCouldNotFetchUserInfo = errors.New("could not fetch user info")
	ErrCouldNotFetchServers  = errors.New("could not fetch servers")
	ErrInvalidProvider       = errors.New("could not fetch servers")
	ErrServersNotFound       = errors.New("could not find any server")
	ErrDownloadTestFailure   = errors.New("could not process download test")
	ErrUploadTestFailure     = errors.New("could not process upload test")
)
