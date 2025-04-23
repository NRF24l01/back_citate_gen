package utils

import (
	"encoding/json"
	"log"
)

// ToJSON converts a Go value to a JSON byte slice.
func ToJSON(value interface{}) []byte {
    jsonData, err := json.Marshal(value)
    if err != nil {
        log.Fatalf("Failed to convert to JSON: %v", err)
    }
    return jsonData
}