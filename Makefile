watch:
	reflex -r '\.go' -s --decoration=none -- sh -c "go run src/main.go"

build-release:
	docker build --target release -t go-rest:latest -f docker/Dockerfile .
