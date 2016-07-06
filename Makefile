BUILD=`date +%FT%T%z`
NAME=api-gateway
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
VERSION=0.1.0

default:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS}

install:
	go install ${LDFLAGS}

clean:
	go clean
	rm -f *.log

delimage:
	docker rmi -f donbstringham/${NAME}
	
dockerize:
	docker build -t donbstringham/${NAME} .

fmt:
	go fmt ./...

get:
	go get ./...