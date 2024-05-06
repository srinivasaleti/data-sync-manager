package mocks

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

// MockConnector mocks a connector and its methods
type MockConnector struct {
	connectors.Connector
	getErr       error
	getKeys      []interface{}
	getResponse  interface{}
	noOfPutCalls int
	noOfGetCalls int
	putErr       error
	putPayload   interface{}
	exists       bool
	keys         []string
	listKeysErr  error
}

func (mock *MockConnector) Get(key string) ([]byte, error) {
	mock.getKeys = append(mock.getKeys, key)
	mock.noOfGetCalls = mock.noOfGetCalls + 1
	byteResponse, err := json.Marshal(mock.getResponse)
	if err != nil {
		return nil, err
	}
	return byteResponse, mock.getErr
}

func (mock *MockConnector) SetGetResponse(data interface{}) {
	mock.getResponse = data
}

func (mock *MockConnector) SetGetErr(err error) {
	mock.getErr = err
}

func (mock *MockConnector) GetShouldBeCalledWith(id string) bool {
	for _, payload := range mock.getKeys {
		fmt.Println(payload)
		if payload == id {
			return true
		}
	}
	return false
}

func (mock *MockConnector) Put(key string, data []byte) error {
	mock.putPayload = data
	mock.noOfPutCalls = mock.noOfPutCalls + 1
	return mock.putErr
}

func (mock *MockConnector) SetPutErr(err error) {
	mock.putErr = err
}

func (mock *MockConnector) PutShouldBeCalledWith(key string, payload interface{}) bool {
	if reflect.DeepEqual(payload, mock.putPayload) == true {
		return true
	}
	return false
}

func (mock *MockConnector) NumberOfPutCalls() int {
	return mock.noOfPutCalls
}

func (mock *MockConnector) NumberOfGetCalls() int {
	return mock.noOfGetCalls
}

func (mock *MockConnector) ToString() string {
	return "s3"
}

func (mock *MockConnector) Exists(key string) bool {
	return mock.exists
}

func (mock *MockConnector) SetExists(exists bool) {
	mock.exists = exists
}

func (mock *MockConnector) ListKeys() ([]string, error) {
	return mock.keys, mock.listKeysErr
}

func (mock *MockConnector) SetListKeys(keys []string) {
	mock.keys = keys
}

func (mock *MockConnector) SetListKeysErr(err error) {
	mock.listKeysErr = err
}

func (mock *MockConnector) Reset() {
	mock.getErr = nil
	mock.getKeys = nil
	mock.getResponse = nil
	mock.noOfGetCalls = 0
	mock.noOfPutCalls = 0
	mock.putErr = nil
	mock.putPayload = nil
	mock.exists = false
}
