package postgres

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-target-connectors/config"
	"math"
	"os"
)

const (
	defaultMaxIdleConnections           = 10
	defaultMaxOpenConnections           = 100
	defaultConnectionMaxLifetimeSeconds = 3600
)

type options struct {
	credentials string
	useProxy    bool

	connection string
	// maxIdleConnections sets the maximum number of connections in the idle connection pool
	maxIdleConnections int
	//maxOpenConnections sets the maximum number of open connections to the database.
	maxOpenConnections int
	// connectionMaxLifetimeSeconds sets the maximum amount of time a connection may be reused.
	connectionMaxLifetimeSeconds int
}

func parseOptions(cfg config.Metadata) (options, error) {
	o := options{}
	o.credentials = cfg.ParseString("credentials", "")
	if o.credentials != "" {
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", o.credentials)
	}
	err := config.MustExistsEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if err != nil {
		return options{}, err
	}

	o.useProxy = cfg.ParseBool("use_proxy", true)

	o.connection, err = cfg.MustParseString("connection")
	if err != nil {
		return options{}, fmt.Errorf("error parsing connection string, %w", err)
	}

	o.maxIdleConnections, err = cfg.ParseIntWithRange("max_idle_connections", defaultMaxIdleConnections, 1, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing max idle connections value, %w", err)
	}
	o.maxOpenConnections, err = cfg.ParseIntWithRange("max_open_connections", defaultMaxOpenConnections, 1, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing max open connections value, %w", err)
	}
	o.connectionMaxLifetimeSeconds, err = cfg.ParseIntWithRange("connection_max_lifetime_seconds", defaultConnectionMaxLifetimeSeconds, 1, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing connection max lifetime seconds value, %w", err)
	}

	return o, nil
}