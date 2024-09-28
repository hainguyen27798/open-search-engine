package settings

type Config struct {
	Server    ServerSettings
	Typesense TypesenseSettings
	MongoDB   MongoDBSettings
}

type ServerSettings struct {
	Port int
	Mode string
}

type TypesenseSettings struct {
	Host   string
	ApiKey string
}

type MongoDBSettings struct {
	URL         string
	Username    string
	Password    string
	MaxPoolSize uint64
}
