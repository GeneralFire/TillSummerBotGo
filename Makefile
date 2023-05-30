mipsle:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat  go build -trimpath -ldflags="-s -w" cmd/TillSummerBot/main.go

upload:
	scp main root@192.168.1.1:/tmp/	

test:
	go test ./...
