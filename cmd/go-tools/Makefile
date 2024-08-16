.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o speedctl main.go

.PHONY: mac
mac:
	GOOS=darwin GOARCH=amd64 go build -o speedctl-mac main.go
