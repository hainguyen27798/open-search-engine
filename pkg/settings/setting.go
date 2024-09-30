package settings

type Config struct {
	Server  ServerSettings
	MongoDB MongoDBSettings
}

type ServerSettings struct {
	Port int
	Mode string
}

type MongoDBSettings struct {
	URL         string
	Username    string
	Password    string
	Database    string
	MaxPoolSize uint64
}
