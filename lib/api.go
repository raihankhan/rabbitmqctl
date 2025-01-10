package lib

type RabbitMQConfig struct {
	RabbitVersion    string            `json:"rabbit_version"`
	RabbitMQVersion  string            `json:"rabbitmq_version"`
	ProductName      string            `json:"product_name"`
	ProductVersion   string            `json:"product_version"`
	Users            []User            `json:"users"`
	Vhosts           []Vhost           `json:"vhosts"`
	Permissions      []Permission      `json:"permissions"`
	TopicPermissions []interface{}     `json:"topic_permissions"`
	Parameters       []interface{}     `json:"parameters"`
	GlobalParameters []GlobalParameter `json:"global_parameters"`
	Policies         []interface{}     `json:"policies"`
	Queues           []Queue           `json:"queues"`
	Exchanges        []interface{}     `json:"exchanges"`
	Bindings         []interface{}     `json:"bindings"`
}

type User struct {
	Name             string            `json:"name"`
	PasswordHash     string            `json:"password_hash"`
	HashingAlgorithm string            `json:"hashing_algorithm"`
	Tags             []string          `json:"tags"`
	Limits           map[string]string `json:"limits"`
}

type Vhost struct {
	Name string `json:"name"`
}

type Permission struct {
	User      string `json:"user"`
	Vhost     string `json:"vhost"`
	Configure string `json:"configure"`
	Write     string `json:"write"`
	Read      string `json:"read"`
}

type GlobalParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Queue struct {
	Name       string            `json:"name"`
	Vhost      string            `json:"vhost"`
	Durable    bool              `json:"durable"`
	AutoDelete bool              `json:"auto_delete"`
	Arguments  map[string]string `json:"arguments"`
}
