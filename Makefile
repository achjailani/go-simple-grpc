proto:
	for f in proto/*/*.proto; do \
		protoc --go_out=. $$f; \
		protoc --go-grpc_out=. $$f; \
		echo compiled: $$f; \
	done

run:
	@go run server.go

build:
	@go build -ldflags="-X 'main.Version=${BUILD_VERSION}'" -v -a -installsuffix -o main .