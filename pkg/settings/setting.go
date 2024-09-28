package settings

type Config struct {
	Server    ServerSettings
	Typesense TypesenseSettings
}

type ServerSettings struct {
	Port int
	Mode string
}

type TypesenseSettings struct {
	Host   string
	ApiKey string
}
