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
	mkdir -p ./build/warmcache/windows
	mkdir -p ./build/warmcache/linux
	mkdir -p ./build/warmcache/macos
	GOOS=darwin GOARCH=amd64 go build -o ./build/warmcache/macos/warmcache $(source)
	GOOS=linux GOARCH=amd64 go build -o ./build/warmcache/linux/warmcache $(source)
	GOOS=windows GOARCH=amd64 go build -o  ./build/warmcache/windows/warmcache.exe $(source)
	cd ./build && tar -czf ./warmcache.tar.gz ./warmcache/
	@echo "publish to gihub: $ hub release create -a ./build/warmcache.tar.gz -m 'v0.X' v0.X"