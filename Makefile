mockprovider:
	mockgen -package mock -destination=api/provider/mock/provider_mock.go github.com/rafdekar/speedtest/api/provider Provider

test:
	go test -v -cover ./...

.PHONY: mockprovider test