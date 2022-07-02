package netflix

import (
	"gopkg.in/ddo/go-fast.v0"
	"sync"
)

// Netflix is provider for www.fast.com website
type Netflix struct {
	dlSpeed float64
}

// Init method is creating new Fast instance, getting urls of servers, measuring download speed and choosing the best result
func (n *Netflix) Init() error {
	fastCom := fast.New()
	err := fastCom.Init()
	if err != nil {
		return err
	}

	urls, err := fastCom.GetUrls()
	if err != nil {
		return err
	}

	KbpsChan := make(chan float64, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	var resultKbps float64
	go func() {
		for p := range KbpsChan {
			// selecting the highest speed of all passes
			if resultKbps < p {
				resultKbps = p
			}
		}

		wg.Done()
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		return err
	}

	wg.Wait()

	n.dlSpeed = resultKbps

	return nil
}

// GetDownloadSpeed returns download speed of a user
func (n *Netflix) GetDownloadSpeed() float64 {
	return n.dlSpeed
}

// GetUploadSpeed returns upload speed of a user
func (n *Netflix) GetUploadSpeed() float64 {
	return n.dlSpeed
}
