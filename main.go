package main

import (
	"github.com/shaksham08/temo_go_mod/pkg/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Info("A walrus appears")

	api.CallApi()
}
