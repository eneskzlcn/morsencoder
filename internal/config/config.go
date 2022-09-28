package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:"server"`
}

/*LoadConfig is a generic function that takes a config structure as a type,
path of the config, and the name of the file in that path, then the format of the file.
Ex:
Suppose that you have a config type like;

type SomeConfig struct {
	url string `mapstructure:"url"`
}

and a file in the .dev/ directory named local.yaml with the following content:


url: "someurl"

Then all you need to call the function to get the configs from the file by:

someConfig , err := LoadConfig[SomeConfig](".dev/", "local", "yaml")

*/
func LoadConfig[T any](path, name, configType string) (config T, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(configType)
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
