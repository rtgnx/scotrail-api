name=scotrail-api

build:
	docker build -t $(name) .

run: build
	docker run -d --name $(name) --net host $(name)

update: build
	docker rm -f $(name) && docker run -d --name $(name) --net host $(name)

