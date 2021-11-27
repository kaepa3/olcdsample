build:
	go build -o display

pi:
	GOOS=linux GOARCH=arm GOARM=6 go build -o display main.go
