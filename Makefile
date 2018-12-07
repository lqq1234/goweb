NAME := goweb
VERSION := v0.1

all: run

clean:
	rm -rf ./GoWeb

image: clean
	docker run --rm -v $(shell pwd):/go/src/$(NAME) -w /go/src/$(NAME) golang:1.8-alpine go build -o $(NAME)
	docker build -t registry.paas/library/$(NAME):$(VERSION) .

run: image
	docker run -dit -v /etc/localtime:/etc/localtime -v /opt/conf:/opt/conf -p 8888:8888 registry.paas/library/$(NAME):$(VERSION)
