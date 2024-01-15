package application

type Configuration struct {
	port string
}

type ConfigurationBuilder struct {
	configuration Configuration
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	return &ConfigurationBuilder{}
}

func (b *ConfigurationBuilder) WithPort(port string) *ConfigurationBuilder {
	if port == "" {
		panic("port cannot be empty")
	}
	b.configuration.port = port
	return b
}

func (b *ConfigurationBuilder) Build() Configuration {
	return b.configuration
}
