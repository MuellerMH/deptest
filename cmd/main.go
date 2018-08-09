package main

import (
	"fmt"

	"../internal/stuff"

	"github.com/sirupsen/logrus"
)

func main() {
	stuff := stuff.Stuff{}
	fmt.Println("hello", stuff)
	logrus.WithFields(logrus.Fields{"layer": "dlv"}).Warnln("CGO_CFLAGS already set, Cgo code could be optimized.")
}
