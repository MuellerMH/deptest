package main

import (
	"fmt"

	"github.com/MuellerMH/deptest/internal/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	stuff := stuff.Stuff{}
	fmt.Println("hello", stuff)
	logrus.WithFields(logrus.Fields{"layer": "dlv"}).Warnln("CGO_CFLAGS already set, Cgo code could be optimized.")
}
