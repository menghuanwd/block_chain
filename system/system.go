package system

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Configuration ...
type Configuration struct {
	PgHost     string `yaml:"pg_host"`
	PgPort     string `yaml:"pg_port"`
	PgDatabase string `yaml:"pg_database"`
	PgUser     string `yaml:"pg_user"`
	PgPassword string `yaml:"pg_password"`
}

var configuration *Configuration

// LoadConfiguration ...
func LoadConfiguration(path, environment string) error {
	if path == "" {
		path = FilePath(environment)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

// GetConfiguration ...
func GetConfiguration() *Configuration {
	return configuration
}

// FilePath ...
func FilePath(environment string) (path string) {
	switch environment {
	case "production":
		path = "config/production.yaml"
	case "staging":
		path = "config/staging.yaml"
	default:
		path = "config/development.yaml"
	}
	return
}
