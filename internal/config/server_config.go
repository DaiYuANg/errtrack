package config

import "strconv"

type ServerConfig struct {
	Port    int
	Prefork bool
}

func (c *ServerConfig) ListenAddress() string {
	return ":" + strconv.Itoa(c.Port)
}
