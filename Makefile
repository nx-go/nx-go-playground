all: build push

build-playground:
	docker build . -f Dockerfile -t nxgo/nx-go-playground:latest

push-playground:
	docker push nxgo/nx-go-playground:latest

build: build-playground
push: push-playground

run:
	docker run -it --rm --name nx-go-playground -p3000:3000 nxgo/nx-go-playground:latest
