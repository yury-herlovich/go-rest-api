watch:
	reflex -r '\.go' -s --decoration=none -- sh -c "go run src/main.go"

docker.watch:
	docker-compose up -d

docker.debug:
	docker-compose run --rm -p1541:1541 -p8080:8080 app dlv debug --listen=0.0.0.0:1541 --headless --api-version=2 ./src/main.go

release-build:
	docker build --target release -t go-rest:latest -f docker/Dockerfile .

release-run:
	docker run -p8080:8080 -d go-rest
