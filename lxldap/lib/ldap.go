package lib

// OpenldapConfig structure to store OpenLDAP config
type OpenldapConfig struct {
	Server   string
	Port     string
	BaseDN   string
	OUusers  string
	OUgroups string
	OUconfig string
	DN       string
	PS       string
}

// SetOpenldapConfigDefaults sets defautl values to OpenLDAP config
func SetOpenldapConfigDefaults() {
	// viper.SetDefault("OpenldapConfig.server", "localhost")
	// viper.SetDefault("OpenldapConfig.port", "9000")
}
