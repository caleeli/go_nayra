package bpmn

type CallbackType = func(event string, body interface{})

var observers map[string][]CallbackType

func init() {
	observers = make(map[string][]CallbackType)
}

// SubscribeAllEvents listens all BPMN events
func SubscribeAllEvents(callback CallbackType) {
	SubscribeEvent("*", callback)
}

// SubscribeEvent Listen to a specific type of BPMN event
func SubscribeEvent(event string, callback CallbackType) {
	if observers[event] == nil {
		observers[event] = make([]CallbackType, 0)
	}
	observers[event] = append(observers[event], callback)
}

// NotifyEvent Nofity a specific type of BPMN event
func NotifyEvent(event string, body interface{}) {
	for i := range observers[event] {
		observers[event][i](event, body)
	}
	for i := range observers["*"] {
		observers["*"][i](event, body)
	}
}
