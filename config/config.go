package config

type Configurations struct {
	Database Database
	Server   Server
}

type Database struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

type Server struct {
	Host string
	Port int
}
