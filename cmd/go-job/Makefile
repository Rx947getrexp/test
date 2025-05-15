.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-job -trimpath main.go

.PHONY: mac
mac:
	GOOS=darwin GOARCH=amd64 go build -o go-job-mac main.go