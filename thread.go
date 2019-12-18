package tsubato

import (
	"encoding/json"
	"fmt"
)

// Thread => []Post
type Thread struct {
	Posts []ThreadPost `json:"posts"`
}

// GetThread returns the thread for a given id on a given board
func GetThread(boardname string, id string) (Thread, error) {
	JSON, err := get(fmt.Sprintf("/%s/thread/%s.json", boardname, id))
	if err != nil {
		return Thread{}, err
	}

	thread := Thread{}
	err = json.Unmarshal(JSON, &thread)
	if err != nil {
		return Thread{}, err
	}

	return thread, nil
}
