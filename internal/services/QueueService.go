package services

import "nayra/internal/bpmn"

type QueueService interface {
	NotifyEvent(event string, body interface{}) (err error)
}

var Queue QueueService

func SetupQueueService(queue QueueService) {
	Queue = queue
	// Subscribe and notify all BPMN events
	//bpmn.SubscribeAllEvents(notifyBpmnEvents)
	bpmn.SubscribeEvent("RUN_SCRIPT", notifyBpmnEvents)
}

func notifyBpmnEvents(event string, body interface{}) {
	Queue.NotifyEvent(event, body)
}
