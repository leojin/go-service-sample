package server

import server "github.com/leojin/go-service-bootstrap-server-echo"

func NewServer() *server.HTTPServerEcho {
	return &server.HTTPServerEcho{
		Name: "admin",
		Port: 8011,
	}
}
