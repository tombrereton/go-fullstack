package server

type Config struct {
	Port          string `envconfig:"PORT" default:"3000"`
	IsDevelopment bool   `envconfig:"IS_DEVELOPMENT" default:"false"`
	TemplateDir   string `envconfig:"TEMPLATES_DIR" default:"templates/"`
	StaticDir     string `envconfig:"STATIC_DIR" default:"static/"`
}
