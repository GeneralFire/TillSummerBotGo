mipsle:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat  go build -trimpath -ldflags="-s -w" cmd/TillSummerBot/main.go
