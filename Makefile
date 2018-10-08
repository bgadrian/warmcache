# Makefile
source := ./src/main.go

pre:
	mkdir -p ./build/
	env GO111MODULE=on go get -d ./src/

run: pre
	go run $(source) --seed $(URL) --debug

build: pre
	go build -o ./build/hotcache $(source)
	@echo "See ./build/hotcache --help"

buildall: pre
	GOOS=darwin GOARCH=amd64 go build -o ./build/hotcache-mac $(source)
	GOOS=linux GOARCH=amd64 go build -o ./build/hotcache $(source)
	GOOS=windows GOARCH=amd64 go build -o  ./build/hotcache.exe $(source)