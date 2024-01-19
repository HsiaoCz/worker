package main

import "fmt"

type Server struct {
	config Config
}

type Config struct {
	ListenAddr string
	ID         string
	Name       string
}

func NewServer(config Config) *Server {
	return &Server{
		config: config,
	}
}

func NewConfig() Config {
	return Config{
		ListenAddr: ":9001",
		ID:         "random_ID",
		Name:       "random_name",
	}
}

func (c Config) WithListenAddr(addr string) Config {
	c.ListenAddr = addr
	return c
}

func (c Config) WithID(ID string) Config {
	c.ID = ID
	return c
}

func (c Config) WithName(name string) Config {
	c.Name = name
	return c
}

func main() {
	config := NewConfig().WithID("Hello").WithListenAddr(":9002").WithName("IDM")
	server := NewServer(config)
	fmt.Println(server)

}
