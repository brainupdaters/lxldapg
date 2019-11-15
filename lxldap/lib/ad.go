package lib

// ADConfig structure to store OpenLDAP config
type ADConfig struct {
	Server   string
	Port     string
	BaseDN   string
	OUusers  string
	OUgroups string
	DN       string
	PS       string
}

// SetADConfigDefaults sets defautl values to OpenLDAP config
func SetADConfigDefaults() {
	// viper.SetDefault("OpenldapConfig.server", "localhost")
	// viper.SetDefault("OpenldapConfig.port", "9000")
}
