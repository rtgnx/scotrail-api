name=scotrail-api
PORT=8049

build:
	docker build -t $(name) .

run: build
	docker run -d --name $(name) --net host $(name)

update: build
	docker rm -f $(name) && docker run -d --name $(name) --net host $(name)

run-dev:
	go build && PORT=$(PORT) ./scotrail-api && rm scotrail-api
