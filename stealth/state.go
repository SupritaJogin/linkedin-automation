package stealth

import (
	"os"
	"gopkg.in/yaml.v3"
)

type State struct {
	RunCount        int    `yaml:"run_count"`
	SentConnections int    `yaml:"sent_connections"`
	SentMessages    int    `yaml:"sent_messages"`
	DailyLimit      int    `yaml:"daily_limit"`
	LastRunDate     string `yaml:"last_run_date"`
}

// Load state from YAML file
func LoadState(path string) (*State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var state State
	err = yaml.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

// Save state to YAML file
func SaveState(path string, state *State) error {
	data, err := yaml.Marshal(state)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
