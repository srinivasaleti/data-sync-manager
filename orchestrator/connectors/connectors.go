package connectors

type Connector interface {
	Get(id string) (interface{}, error)
	Put(v interface{}) error
	ToString() string
}
