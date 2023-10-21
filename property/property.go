package property

import (
	"context"
	"log"

	configs "github.com/centraldigital/cfw-core-lib/pkg/configuration/backoffice-property"
	"github.com/kelseyhightower/envconfig"
)

func Init(ctx context.Context) {
	configs.InitServerProperties()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() config {
	return cfg
}

var cfg config

type config struct {
	Service serviceProperties
	DB      postgres
}

type serviceProperties struct {
	configs.ServerProperties
}

type postgres struct {
	PostgresConnectionUri string `envconfig:"POSTGRES_CONNECTION_URI"`
	Test                  int    `envconfig:"TEST"`
}
