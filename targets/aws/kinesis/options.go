package kenisis

import (
	"fmt"

	"github.com/kubemq-hub/kubemq-targets/config"
)

type options struct {
	awsKey       string
	awsSecretKey string
	region       string
	token        string
}

func parseOptions(cfg config.Spec) (options, error) {
	o := options{}
	var err error
	o.awsKey, err = cfg.MustParseString("aws_key")
	if err != nil {
		return options{}, fmt.Errorf("error awsKey , %w", err)
	}

	o.awsSecretKey, err = cfg.MustParseString("aws_secret_key")
	if err != nil {
		return options{}, fmt.Errorf("error awsSecretKey , %w", err)
	}

	o.region, err = cfg.MustParseString("region")
	if err != nil {
		return options{}, fmt.Errorf("error region , %w", err)
	}

	o.token = cfg.ParseString("token", DefaultToken)

	return o, nil
}
