lint:
	@golangci-lint run ./...

mocks:
	mockgen -source=internal/infrastructure/router/router.go -destination=internal/infrastructure/server/router_mock_test.go -package=server
	mockgen -source=internal/usecase/mocker/usecase.go -destination=internal/infrastructure/server/usecase_mock_test.go -package=server
	mockgen -source=internal/usecase/mocker/services.go -destination=internal/usecase/mocker/services_mock_test.go -package=mocker

test:
	@go test -coverpkg=./internal/... -coverprofile=cover.out ./internal/...

cover: test
	@go tool cover -func cover.out

build:
	@echo "  >  Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mocker cmd/mocker/main.go

clean:
	@go clean
