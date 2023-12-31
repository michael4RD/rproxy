build:
	go build -o bin/main main.go
	go build -o bin/target target_server.go

run:
	go run target_server.go &
	go run main.go

clean:
	rm -rf bin

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

all: build
