package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
	"verve/database"
)

var uniqueIDs = sync.Map{}

func StartTicker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		logAndResetUniqueCount()
	}
}

func logAndResetUniqueCount() {
	// Count unique keys in Redis

	keys, err := database.RedisClient.Keys("*").Result()
	if err != nil {
		log.Printf("Failed to fetch keys from Redis: %v", err)
		return
	}

	uniqueCount := len(keys)
	log.Printf("Unique requests in the last minute: %d", uniqueCount)

	// Clear Redis keys
	for _, key := range keys {
		database.RedisClient.Del(key)
	}
}

func MakeHTTPRequest(endpoint string) {
	// Build payload
	payload := map[string]interface{}{
		"unique_count": len(getCurrentUniqueIDs()),
	}

	jsonData, _ := json.Marshal(payload)

	// Make POST request

	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Printf("Failed to send POST request: %v", err)
		return
	}

	log.Printf("HTTP response status: %s", resp.Status)
}

func getCurrentUniqueIDs() []int {
	var ids []int
	uniqueIDs.Range(func(key, value interface{}) bool {
		ids = append(ids, key.(int))
		return true
	})
	return ids
}
