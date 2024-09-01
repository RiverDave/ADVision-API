package config

type Config struct {
	Oakey string
	// debug / prod?
	Env string
}

func InitConfig(apikey string, env string) *Config {
	// Horseshit error handling but it's a prototype right?
	if apikey == "" {
		panic("API key not set. Exiting...")
	}

	if env == "" {
		panic("Environment not set. Exiting...")
	}

	return &Config{
		Oakey: apikey,
		Env:   env,
	}
}

// Unsure on how safe this is
func (c *Config) OaKey() string {
	return c.Oakey
}

func (c *Config) Environment() string {
	return c.Env
}
