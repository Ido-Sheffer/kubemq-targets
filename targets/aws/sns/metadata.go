package sns

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-targets/types"
)

type metadata struct {
	method string
	topic  string

	endPoint           string
	protocol           string
	returnSubscription bool

	message           string
	phoneNumber       string
	subject           string
	targetArn         string
	messageAttributes map[string]string
}

var methodsMap = map[string]string{
	"list_topics":              "list_topics",
	"listSubscriptions":        "listSubscriptions",
	"listSubscriptionsByTopic": "listSubscriptionsByTopic",
	"create_topic":             "create_topic",
	"subscribe":                "subscribe",
	"send_message":             "send_message",
	"delete_topic":             "delete_topic",
}

func getValidMethodTypes() string {
	s := "invalid method type, method type should be one of the following:"
	for k := range methodsMap {
		s = fmt.Sprintf("%s :%s,", s, k)
	}
	return s
}

func parseMetadata(meta types.Metadata) (metadata, error) {
	m := metadata{}
	var err error
	m.method, err = meta.ParseStringMap("method", methodsMap)
	if err != nil {
		return metadata{}, fmt.Errorf(getValidMethodTypes())
	}
	if m.method != "list_topics" || m.method != "listSubscriptions" {
		m.topic, err = meta.MustParseString("topic")
		if err != nil {
			return metadata{}, fmt.Errorf("error parsing topic, %w", err)
		}
		if m.method == "subscribe"{
			m.endPoint, err = meta.MustParseString("end_point")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing end_point, %w", err)
			}
			m.protocol, err = meta.MustParseString("protocol")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing protocol, %w", err)
			}
			m.returnSubscription, err = meta.MustParseBool("return_subscription")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing return_subscription, %w", err)
			}
		}else if m.method == "send_message"{
			m.message, err = meta.MustParseString("message")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing message, %w", err)
			}
			m.phoneNumber, err = meta.MustParseString("phone_number")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing phone_number, %w", err)
			}
			m.subject, err = meta.MustParseString("subject")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing subject, %w", err)
			}
			m.targetArn, err = meta.MustParseString("target_arn")
			if err != nil {
				return metadata{}, fmt.Errorf("error parsing target_arn, %w", err)
			}
		}
	}
	return m, nil
}
