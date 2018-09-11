package main

import (
	"time"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true, TimestampFormat: time.Kitchen})
	log.SetLevel(log.DebugLevel)
	n := negroni.New()
	n.Use(negroni.NewRecovery())

	n.UseHandler(NewRouter())
	n.Run(":4444")
}
