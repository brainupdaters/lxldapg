package lib

import "github.com/spf13/viper"

// ApiserverConfig structure to store API config
type ApiserverConfig struct {
	Port string
	Cert string
	Key  string
}

// SetApiserverConfigDefaults sets defautl values to API server config
func SetApiserverConfigDefaults() {
	viper.SetDefault("ApiserverConfig.port", "8080")
}
