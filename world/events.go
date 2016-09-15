package world

import (
	"sync"
)

type Event interface {
	GetEventType() string
}

var initOnce = &sync.Once{}
var eventHistory []Event = make([]Event, 0)
var eventLock = &sync.RWMutex{}

var eventChannel chan Event = make(chan Event, 100)

func init()  {
	initOnce.Do(func() {
		go processEvents(eventChannel)
	})
}

func logEvent(event Event) {
	eventChannel<- event
}

func processEvents(ec chan Event) {
	for true {
		e, ok := <-ec
		if !ok {
			break;
		}

		eventLock.Lock()
		eventHistory = append(eventHistory, e)
		eventLock.Unlock()
	}
}

func GetEventHistory() []Event {
	eventLock.RLock()
	out := make([]Event, len(eventHistory))
	copy(out, eventHistory)
	eventLock.RUnlock()

	return out
}