package main

import (
	"flag"
	"time"

	v1 "github.com/scottd018/payments-api/api/v1"
	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/seeds"

	log "github.com/sirupsen/logrus"
)

var (
	address         = flag.String("addr", ":8080", "Local address for the HTTP API")
	logLevel        = flag.String("log-level", "INFO", "Log-level (ERROR|WARN|INFO|DEBUG|TRACE)")
	initialSeedFile = flag.String("initial-seed-file", "", "Run one-time seeds passing path to a valid JSON seed file")
)

func configureLogging() error {
	l, err := log.ParseLevel(*logLevel)
	if err != nil {
		return err
	}

	log.SetLevel(l)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
	})

	return nil
}

func main() {
	flag.Parse()
	if err := configureLogging(); err != nil {
		log.Fatal(err)
	}

	err := dbs.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	if initialSeedFile != nil && *initialSeedFile != "" {
		if err = seeds.RunSeeds(*initialSeedFile); err != nil {
			log.Fatal(err)
		}
	}

	api := v1.NewAPI()
	if err = api.Serve(*address); err != nil {
		log.Fatal(err)
	}
}
