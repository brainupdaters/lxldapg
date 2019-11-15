package lib

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// LxldapConfig lxldap configuration data
type LxldapConfig struct {
	Openldap  OpenldapConfig  `mapstructure:"openldap"`
	Ad        ADConfig        `mapstructure:"ad"`
	Logging   LoggingConfig   `mapstructure:"logging"`
	Apiserver ApiserverConfig `mapstructure:"apiserver"`
}

// Config to storage
var Config *LxldapConfig

// SetConfigDefaults Set config to default values
func SetConfigDefaults() {
	SetOpenldapConfigDefaults()
	SetADConfigDefaults()
	SetLoggingConfigDefaults("lxldap")
	SetApiserverConfigDefaults()
}

// InitConfig initialize the configuration
func InitConfig(c string) {

	if c != "" {
		// Use config file from the flag.
		viper.SetConfigFile(c)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lxldap" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./config/")
		viper.SetConfigName(".lxldap")
	}

	SetConfigDefaults()

	// Enable environment variables
	// ex.: LXLDAP_LDAP_PORT=8000
	viper.SetEnvPrefix("LXLDAP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to find config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		panic("Unable to unmarshal config")
	}
}
