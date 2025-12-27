package storage

import (
	"encoding/json"
	"os"
	"time"
)

type State struct {
	SentConnections []string  `json:"sent_connections"`
	SentMessages    []string  `json:"sent_messages"`
	LastRun         time.Time `json:"last_run"`
}

func LoadState(path string) (*State, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &State{
			SentConnections: []string{},
			SentMessages:    []string{},
			LastRun:         time.Now(),
		}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var state State
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func SaveState(path string, state *State) error {
	state.LastRun = time.Now()

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
