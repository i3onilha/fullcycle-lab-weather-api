.PHONY: build run test docker-build docker-run clean

build:
	go build -o bin/weather-api main.go

run:
	go run main.go

test:
	go test -v ./...

docker-build:
	docker build -t weather-api .

docker-run:
	docker-compose up --build

clean:
	rm -rf bin/
