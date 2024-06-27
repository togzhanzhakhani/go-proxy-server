.PHONY: build run docker-build docker-run

build:
	go build -o go-proxy-server .

run: build
	./go-proxy-server

docker-build:
	docker build -t go-proxy-server .

docker-run:
	docker-compose up --build
