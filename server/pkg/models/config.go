package models

type Config struct {
	Server   *Server
	Database *Database
}

type Server struct {
	Name string
	Port string
}

type Database struct {
	Type     string
	Postgres *Postgres
}

type Postgres struct {
	Username     string `json:"-"`
	Password     string `json:"-"`
	Port         int
	DatabaseName string
	Host         string
	Schema       string
}
