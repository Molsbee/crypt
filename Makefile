build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

build-run: build
	./crypt
