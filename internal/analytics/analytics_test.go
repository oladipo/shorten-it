package analytics

import (
	"testing"
	"time"
)

func TestRecordAndGetEvents(t *testing.T) {
	shortcode := "test123"
	event := Event{
		Timestamp: time.Now(),
		IP:        "1.2.3.4",
		Country:   "Wonderland",
		City:      "Rabbit Hole",
		Referrer:  "https://example.com",
		UserAgent: "TestAgent/1.0",
	}

	Record(shortcode, event)
	got := GetEvents(shortcode)

	if len(got) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(got))
	}
	if got[0].IP != event.IP || got[0].Country != event.Country || got[0].City != event.City {
		t.Errorf("Event fields not stored correctly: got %+v, want %+v", got[0], event)
	}
}
