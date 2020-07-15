package minio

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-target-connectors/config"
)

type options struct {
	endpoint        string
	useSSL          bool
	accessKeyId     string
	secretAccessKey string
}

func parseOptions(cfg config.Metadata) (options, error) {
	o := options{}
	var err error
	o.endpoint, err = cfg.MustParseString("endpoint")
	if err != nil {
		return options{}, fmt.Errorf("error parsing endpoint, %w", err)
	}
	o.useSSL = cfg.ParseBool("use_ssl", true)
	o.accessKeyId, err = cfg.MustParseEnv("access_key_id", "ACCESS_KEY_ID", "")
	if err != nil {
		return options{}, fmt.Errorf("error parsing access key id, %w", err)
	}

	o.secretAccessKey, err = cfg.MustParseEnv("secret_access_key", "SECRET_ACCESS_KEY", "")
	if err != nil {
		return options{}, fmt.Errorf("error parsing secret access key, %w", err)
	}
	return o, nil
}