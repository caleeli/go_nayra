package services

import (
	"encoding/json"
	"log"
)

type tLogQueueService struct{}

func LogQueue() (QueueService, error) {
	service := &tLogQueueService{}
	return service, nil
}

func (service *tLogQueueService) NotifyEvent(event string, body interface{}) (err error) {
	log.Println("EVENT:", event)
	msg, err := json.Marshal(body)
	log.Println("      ", string(msg))
	return
}
