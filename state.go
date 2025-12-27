package stealth

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// State represents saved state for LinkedIn automation
type State struct {
	LoggedIn        bool `yaml:"logged_in"`
	SentConnections int  `yaml:"sent_connections"`
	SentMessages    int  `yaml:"sent_messages"`
}

// LoadState reads state from a YAML file
func LoadState(filename string) (*State, error) {
	data, err := os.ReadFile(filename)
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

// SaveState writes state to a YAML file
func SaveState(filename string, state *State) error {
	data, err := yaml.Marshal(state)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
