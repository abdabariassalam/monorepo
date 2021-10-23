package config

type (
	// Config .
	Config struct {
		// JWT.
		JWT struct {
			SecretKey string `yaml:"secret_key"`
		} `ymal:"jwt"`

		// Resource.
		Resource struct {
			Url string `yaml:"url"`
		} `ymal:"resource"`
	}
)
