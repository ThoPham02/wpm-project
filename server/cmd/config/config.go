package config


type Config struct {
	Http     HttpServer 
	Database Database
}

type HttpServer struct {
	Path string 
	Port string
}

type Database struct {
	Driver string
	Source string
}
