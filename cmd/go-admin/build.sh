go build -o ./go-admin main.go && cp ./go-admin /wwwroot/go/go-admin/backup/ && ps -ef | grep go-admin | grep -v 'grep' | awk '{print $2}' | xargs kill && cd /wwwroot/go/go-admin/ && cp -rf backup/go-admin ./ && ./restart.sh


@linux
  GOOS=linux GOARCH=amd64 go build -o go-api main.go
