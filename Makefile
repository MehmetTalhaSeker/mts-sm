# Run the linter on the specified path
lint:
	go mod tidy
	go vet ./...
	go fmt ./...
	gci write -s standard -s default -s "prefix(github.com/MehmetTalhaSeker/mts-sm)" .
	gofumpt -l -w .
	wsl -fix ./...
	golangci-lint run $(p)
.PHONY: lint


# Install linter dependencies
lint-dep:
	go install github.com/daixiang0/gci@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/bombsimon/wsl/v4/cmd...@master