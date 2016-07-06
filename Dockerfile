FROM alpine:latest
MAINTAINER Don B. Stringham <donbstringham@gmail.com> @donbstringham

ADD api-gateway /go/bin/main
ENTRYPOINT /go/bin/main
