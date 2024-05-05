package connectors

type Config struct {
	Type   string
	Config interface{}
}

type Connector interface {
	Get(key string) ([]byte, error)
	Put(key string, data []byte) error
	ToString() string
}
