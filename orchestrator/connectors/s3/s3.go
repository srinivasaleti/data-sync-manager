package s3

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

type S3Connector struct {
	connectors.Connector
}

func (s *S3Connector) Get(id string) (interface{}, error) {
	return "hello", nil
}

func (s *S3Connector) Put(data interface{}) error {
	return nil
}

func (s *S3Connector) ToString() string {
	return "s3"
}

func New() *S3Connector {
	return &S3Connector{}
}
