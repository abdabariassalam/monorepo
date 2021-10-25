package config

type (
	// Config .
	Config struct {
		// Base
		Base struct {
			Port string `yaml:"port"`
		} `yaml:"base"`

		// JWT.
		JWT struct {
			SecretKey string `yaml:"secret_key"`
		} `ymal:"jwt"`

		// Resource.
		Resource struct {
			Url string `yaml:"url"`
		} `ymal:"resource"`

		// Currency
		CurrencyConverter struct {
			BaseUrl string `yaml:"base_url"`
			ApiKey  string `yaml:"api_key"`
		} `yaml:"currency_converter"`
	}
)
