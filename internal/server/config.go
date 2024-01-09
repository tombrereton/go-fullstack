package server

type Config struct {
	IsDevelopment bool   `envconfig:"IS_DEVELOPMENT" default:"false"`
	TemplatesPath string `envconfig:"TEMPLATES_PATH" default:"templates/"`
}
