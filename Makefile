watch:
	reflex -r '\.go' -s --decoration=none -- sh -c "go run src/main.go"

docker-watch:
	docker-compose up -d

release-build:
	docker build --target release -t go-rest:latest -f docker/Dockerfile .

release-run:
	docker run -p8080:8080 -d go-rest
