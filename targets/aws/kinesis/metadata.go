package kenisis

import (
	"github.com/kubemq-hub/kubemq-targets/types"
)

type metadata struct {
	method        string
	stream        string
	partition_key string
	shard_count   int
}

func parseMetadata(meta types.Metadata, opts options) (metadata, error) {
	m := metadata{}
	var err error
	m.stream, err = meta.MustParseString("stream")
	if err != nil {
		return m, err
	}

	m.method = meta.ParseString("method", "put")

	if m.method == "create" {
		m.shard_count = meta.ParseInt("shard_count", 1)
	} else {
		m.partition_key, err = meta.MustParseString("partition_key")
		if err != nil {
			return m, err
		}
	}

	return m, nil
}
