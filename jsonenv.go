package jsonenv

import (
	"os"

	"encoding/json"

	"github.com/pkg/errors"
)

// jsonenv.Extract gets environment name and parse JSON, and set to environment key-value.
func Extract(envName string) error {
	if envName == "" {
		return errors.New("Environment key name is empty")
	}

	env := os.Getenv(envName)
	if env == "" {
		return errors.New("Environment value is empty for: " + envName)
	}

	var kv map[string]string
	if err := json.Unmarshal([]byte(env), &kv); err != nil {
		return errors.Wrap(err, "Failed to unmarshal env JSON")
	}

	for key, val := range kv {
		os.Setenv(key, val)
	}

	return nil
}
