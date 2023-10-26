package utils

// GetConfigPath for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "development" {
		return "./config/config-dev"
	} else if configPath == "heroku" {
		return "./config/config-heroku"
	} else if configPath == "production" {
		return "./config/config-prod"
	} else {
		return "./config/cfg"
	}
}
