package kenisis

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/kinesis"
	awskinesis "github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/kubemq-hub/kubemq-targets/config"
	"github.com/kubemq-hub/kubemq-targets/types"
)

type Client struct {
	name   string
	opts   options
	client *awskinesis.Kinesis
}

func New() *Client {
	return &Client{}

}
func (c *Client) Name() string {
	return c.name
}

func (c *Client) Init(ctx context.Context, cfg config.Spec) error {
	c.name = cfg.Name
	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(c.opts.region),
		Credentials: credentials.NewStaticCredentials(c.opts.awsKey, c.opts.awsSecretKey, c.opts.token),
	})
	if err != nil {
		return err
	}

	c.client = awskinesis.New(sess)

	return nil
}

func (c *Client) Do(ctx context.Context, request *types.Request) (*types.Response, error) {
	m, err := parseMetadata(request.Metadata, c.opts)
	if err != nil {
		return nil, err
	}

	switch m.method {
	case "create":
		o, err := c.client.CreateStreamWithContext(ctx, &kinesis.CreateStreamInput{
			ShardCount: aws.Int64(int64(m.shard_count)),
			StreamName: &m.stream,
		})
		if err != nil {
			return nil, err
		}
		if err := c.client.WaitUntilStreamExistsWithContext(ctx, &kinesis.DescribeStreamInput{StreamName: &m.stream}); err != nil {
			panic(err)
		}

		s, err := c.client.DescribeStreamWithContext(ctx, &kinesis.DescribeStreamInput{StreamName: &m.stream})
		if err != nil {
			return nil, err
		}
		return types.NewResponse().
			SetMetadataKeyValue("CreateStreamOutput", o.String()).
			SetMetadataKeyValue("DescribeStreamOutput", s.String()), nil
	case "put":
		p, err := c.client.PutRecordWithContext(ctx, &kinesis.PutRecordInput{
			Data:         request.Data,
			StreamName:   &m.stream,
			PartitionKey: &m.partition_key,
		})
		if err != nil {
			return nil, err
		}
		return types.NewResponse().
			SetMetadataKeyValue("PutRecordOutput", p.String()), nil
	}
	return nil, fmt.Errorf("not implemented method %s", m.method)

}
