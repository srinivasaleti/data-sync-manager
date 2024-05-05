package mocks

import (
	"encoding/json"
	"reflect"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

// MockConnector mocks a connector and its methods
type MockConnector struct {
	connectors.Connector
	getErr       error
	getPayload   interface{}
	getResponse  interface{}
	noOfPutCalls int
	noOfGetCalls int
	putErr       error
	putPayload   interface{}
}

func (s *MockConnector) Get(key string) ([]byte, error) {
	s.getPayload = key
	s.noOfGetCalls = s.noOfGetCalls + 1
	byteResponse, err := json.Marshal(s.getResponse)
	if err != nil {
		return nil, err
	}
	return byteResponse, s.getErr
}

func (s *MockConnector) SetGetResponse(data interface{}) {
	s.getResponse = data
}

func (s *MockConnector) SetGetErr(err error) {
	s.getErr = err
}

func (s *MockConnector) GetShouldBeCalledWith(id string) bool {
	if s.getPayload == id {
		return true
	}
	return false
}

func (s *MockConnector) Put(key string, data []byte) error {
	s.putPayload = data
	s.noOfPutCalls = s.noOfPutCalls + 1
	return s.putErr
}

func (s *MockConnector) SetPutErr(err error) {
	s.putErr = err
}

func (s *MockConnector) PutShouldBeCalledWith(key string, payload interface{}) bool {
	if reflect.DeepEqual(payload, s.putPayload) == true {
		return true
	}
	return false
}

func (s *MockConnector) NumberOfPutCalls() int {
	return s.noOfPutCalls
}

func (s *MockConnector) NumberOfGetCalls() int {
	return s.noOfGetCalls
}

func (s *MockConnector) ToString() string {
	return "s3"
}

func (s *MockConnector) Reset() {
	s.getErr = nil
	s.getPayload = nil
	s.getResponse = nil
	s.noOfGetCalls = 0
	s.noOfPutCalls = 0
	s.putErr = nil
	s.putPayload = nil
}
