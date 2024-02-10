package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/server"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logrus.SetLevel(ll)
}

func main() {
	server.StartServer()
}
