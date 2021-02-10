all: build push

build-playground:
	docker build . -f Dockerfile -t nxgo/nx-go-playground:latest

push-playground:
	docker push nxgo/nx-go-playground:latest

build: build-playground
push: push-playground

run:
	docker run -it --rm --name nx-go-playground -p3000:3000 nxgo/nx-go-playground:latest


prisma-apply: prisma-format prisma-push

prisma-format:
	go run github.com/prisma/prisma-client-go format

prisma-push:
	go run github.com/prisma/prisma-client-go db push --preview-feature
