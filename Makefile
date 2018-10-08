# Makefile
source := main.go

pre:
	mkdir -p ./build/
	env GO111MODULE=on go get -d ./

run: pre
	go run $(source) --seed $(URL) --debug

build: pre
	go build -o ./build/warmcache $(source)
	@echo "See ./build/warmcache --help"

buildall: pre
	GOOS=darwin GOARCH=amd64 go build -o ./build/warmcache-mac $(source)
	GOOS=linux GOARCH=amd64 go build -o ./build/warmcache $(source)
	GOOS=windows GOARCH=amd64 go build -o  ./build/warmcache.exe $(source)