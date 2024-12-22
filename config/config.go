// Viper library is used here to extract environment variables.
// This is done because it is wildly popular and provides convenient, extensible functionality
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

func validatePort(port int) bool {
	return port >= 0 && port <= 65535
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(fmt.Sprintf("%s/.env", path))

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	if !validatePort(config.Port) {
		err = fmt.Errorf("invalid port: %d", config.Port)
		return
	}

	return
}
