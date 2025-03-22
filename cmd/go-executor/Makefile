.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o executor -trimpath main.go

.PHONY: mac
mac:
	GOOS=darwin GOARCH=amd64 go build -o executor-mac main.go

# ps -ef | grep executor | grep -v 'grep' | awk '{print $2}' | xargs kill && cd /NodeService/ && cp -rf backup/executor ./
# && ./restart.sh