package main

import (
	"net/http"

	"github.com/niclasgeiger/concourse-test/pkg/handler"
	"github.com/sirupsen/logrus"
)

const (
	port = ":8888"
)

func main() {
	http.HandleFunc("/", handler.HandleRandomCalculation)
	logrus.Fatal(http.ListenAndServe(port, nil))
}
