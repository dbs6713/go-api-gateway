package main

import (
	"github.com/bfirsh/go-dcgi"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types/container"
	"net/http"
)

var (
	Version string
	Build   string
)

func main() {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.23", nil, nil)
	if err != nil {
		panic(err)
	}

	hostConfig := &container.HostConfig{
		NetworkMode: "apigateway_default",
		Binds:       []string{"/var/run/docker.sock:/var/run/docker.sock"},
	}

	http.Handle("/", &dcgi.Handler{
		Image:      "donbstringham/serverless-docs",
		Client:     cli,
		HostConfig: hostConfig,
		Root:       "/",
	})

	http.Handle("/hello/", &dcgi.Handler{
		Image:      "donbstringham/serverless-hello",
		Client:     cli,
		HostConfig: hostConfig,
		Root:       "/hello",
	})

	http.ListenAndServe(":80", nil)
}
