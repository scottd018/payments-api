// @title           Payments API
// @version         1.0
// @description     This is the payments API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
