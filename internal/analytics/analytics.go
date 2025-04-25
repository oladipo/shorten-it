package analytics

import (
	"sync"
	"time"
)

type Event struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Referrer  string    `json:"referrer"`
	UserAgent string    `json:"user_agent"`
}

var (
	data = make(map[string][]Event)
	mu   sync.Mutex
)

// Record an analytics event for a shortcode
func Record(shortcode string, e Event) {
	mu.Lock()
	defer mu.Unlock()
	data[shortcode] = append(data[shortcode], e)
}

// GetEvents returns all analytics events for a shortcode
func GetEvents(shortcode string) []Event {
	mu.Lock()
	defer mu.Unlock()
	return append([]Event(nil), data[shortcode]...)
}
