package utils

// GetConfigPath for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "dev" {
		return "./config/config-dev"
	} else if configPath == "heroku" {
		return "./config/config-heroku"
	} else if configPath == "prod" {
		return "./config/config-prod"
	} else {
		return "./config/cfg"
	}
}
