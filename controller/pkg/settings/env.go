package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go-http/pkg/convert"
)

// env config file struct
type EnvConfig struct {
	Backend
	Database
}

// env backend settings
type Backend struct {
	FrontendURL  string
	HostURL      string
	Port         int
	Host         string
	APIPath      string
	RunningInDev bool
}

// env database settings
type Database struct {
	Host     string
	User     string
	Pass     string
	Name     string
	SSL      string
	Timezone string
	Port     int
}

// NewEnvConfig loads the .env file with the specified path, and
// retuns the EnvConfig struct with the loaded values.
func NewEnvConfig(path string) (*EnvConfig, error) {
	if err := godotenv.Load(path); err != nil {
		return nil, fmt.Errorf("error loading .env file, %v", err)
	}

	backendPort, err := convert.StringToInt(os.Getenv("GO_BACKEND_PORT"))
	if err != nil {
		return nil, err
	}

	backendHost := os.Getenv("GO_BACKEND_HOST")

	runningInDev, err := convert.StringToBool(os.Getenv("RUNNING_IN_DEV"))
	if err != nil {
		return nil, err
	}

	dbPort, err := convert.StringToInt(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	env := EnvConfig{
		Backend{
			HostURL:      os.Getenv("HOST_URL"),
			Port:         backendPort,
			Host:         backendHost,
			APIPath:      os.Getenv("GO_BACKEND_API_PATH"),
			RunningInDev: runningInDev,
		},
		Database{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
			SSL:      os.Getenv("DB_SSL"),
			Timezone: os.Getenv("DB_TZ"),
			Port:     dbPort,
		},
	}

	return &env, nil
}
