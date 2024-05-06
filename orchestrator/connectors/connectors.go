package connectors

type Config struct {
	Type   string            `yaml:"type"`
	Config map[string]string `yaml:"config"`
}

type Connector interface {
	ListKeys(callback func(keys []string)) ([]string, error)
	Get(key string) ([]byte, error)
	Put(key string, data []byte) error
	Exists(key string) bool
	ToString() string
}
