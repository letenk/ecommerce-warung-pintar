package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by viper from a config file or environtment variables.
type Config struct {
	ENV        string `mapstructure:"ENV"`
	AppPort    string `mapstructure:"APP_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBNameTest string `mapstructure:"DB_NAME_TEST"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, nameApp string) (config Config, err error) {
	// Config viper
	viper.AddConfigPath(path)
	viper.SetConfigName(nameApp)
	viper.SetConfigType("env")

	// Check all variable in env
	viper.AutomaticEnv()

	// Find and read variable the config file
	err = viper.ReadInConfig()
	// If error
	if err != nil {
		return
	}

	// Insert value config into object viper
	err = viper.Unmarshal(&config)
	return
}
