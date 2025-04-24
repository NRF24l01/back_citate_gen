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

// FromJSONToStringArray converts a JSON byte slice to a slice of strings.
func FromJSONToStringArray(jsonData []byte) []string {
    var result []string
    err := json.Unmarshal(jsonData, &result)
    if err != nil {
        log.Fatalf("Failed to convert from JSON: %v", err)
    }
    return result
}