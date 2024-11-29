package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"verve/database"
	"verve/kafka"
	"verve/utils"

	"github.com/gin-gonic/gin"
)

var topic = "test"

type kafkaStruct struct {
	ID      int
	Message string
}

func HandleRequest(c *gin.Context) {
	// Parse query parameters
	idParam := c.Query("id")
	endpoint := c.Query("endpoint")

	// Validate ID parameter
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "failed")
		return
	}

	// Store the unique ID in Redis
	if err := database.RedisClient.SetNX(strconv.Itoa(id), true, time.Minute).Err(); err != nil {
		log.Printf("Failed to store ID in Redis: %v", err)
		c.String(http.StatusInternalServerError, "failed")
		return
	}

	// Make HTTP call if endpoint is provided
	if endpoint != "" {
		utils.MakeHTTPRequest(endpoint)
	}

	msg := fmt.Sprint("unique id in kafka %d", id)
	kafkaMsg := &kafkaStruct{id, msg}
	byteMsg, err := json.Marshal(kafkaMsg)

	kafka.PushMsgToQueue("topic", byteMsg)
	c.String(http.StatusOK, "ok")
}
