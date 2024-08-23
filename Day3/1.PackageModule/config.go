package main

type Config struct {
	ServiceName string // Contains service name
	Description string // describes the service
}

func (c Config) GetServiceName() string {
	return c.ServiceName
}

func (c Config) GetDescription() string {
	return c.Description
}
