mockprovider:
	mockgen -package mock -destination=api/provider/mock/provider_mock.go github.com/rafdekar/speedtest/api/provider Provider

test:
	go test -v -cover ./...

benchmark:
	go test -bench=. -count=5

.PHONY: mockprovider test benchmark