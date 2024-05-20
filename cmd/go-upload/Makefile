.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-upload main.go

.PHONY: mac
mac:
	GOOS=darwin GOARCH=amd64 go build -o go-upload-mac main.go

# ps -ef | grep go-upload | grep -v 'grep' | awk '{print $2}' | xargs kill && cd /wwwroot/go/go-upload/ && cp -rf backup/go-upload ./ && ./restart.sh
