COMMIT = $(shell git rev-parse HEAD | head -c 8)
TIMESTAMP=$(shell TZ='UTC' date '+%Y%m%dt%H%M')

tag := "$(TIMESTAMP)-$(COMMIT)"

get:
	go get -t -v

build: get
	go vet
	GOOS=linux go  build -ldflags "-X 'main.GitCommit=$(COMMIT)' -X 'main.BuildTime=$(TIMESTAMP)'" -o bin/httpecho .
	GOOS=darwin go build -ldflags "-X 'main.GitCommit=$(COMMIT)' -X 'main.BuildTime=$(TIMESTAMP)'" -o bin/httpecho_osx .

build-docker: build
	docker build -t totomz84/httpecho:$(COMMIT) .
	docker tag totomz84/httpecho:$(COMMIT) totomz84/httpecho:latest 

docker-push: build-docker
	docker push --all-tags totomz84/httpecho