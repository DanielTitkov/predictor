package configs

import "fmt"

type ServerConfig struct {
	Port    int
	TLSPort int `yaml:"tlsPort"`
	Host    string
}

func (s *ServerConfig) GetAddress(tls bool) string {
	port := s.Port
	if tls {
		port = s.TLSPort
	}
	return fmt.Sprintf("%s:%d", s.Host, port)
}
