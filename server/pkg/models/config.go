package models

type Config struct {
	Server *Server
}

type Server struct {
	Name string
	Port string
}
