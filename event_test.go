package gbf_test

import (
	"testing"

	"github.com/KuroiKitsu/go-gbf"
)

func TestCurrentEvents(t *testing.T) {
	events, err := gbf.CurrentEvents()
	if err != nil {
		t.Fatal(err)
	}

	if len(events) == 0 {
		t.Skip("no events found")
	}

	for i, event := range events {
		if event == nil {
			t.Errorf("nil event %d", i)
			continue
		}
		if event.Title == "" {
			t.Errorf("empty title on event %d", i)
		}
		if event.URL == "" {
			t.Errorf("empty url on event %d", i)
		}
	}
}

func TestUpcomingEvents(t *testing.T) {
	events, err := gbf.UpcomingEvents()
	if err != nil {
		t.Fatal(err)
	}

	if len(events) == 0 {
		t.Skip("no events found")
	}

	for i, event := range events {
		if event == nil {
			t.Errorf("nil event %d", i)
			continue
		}
		if event.Title == "" {
			t.Errorf("empty title on event %d", i)
		}
	}
}

func TestEventDetailsURL(t *testing.T) {
	var (
		err error
	)

	details, err := gbf.EventDetailsURL("https://gbf.wiki/Alchemist_Astray")
	if err != nil {
		t.Fatal(err)
	}

	if details == nil {
		t.Fatal("nil result")
	}

	if details.StartsAt.IsZero() {
		t.Error("zero StartsAt")
	}
	if details.EndsAt.IsZero() {
		t.Error("zero EndsAt")
	}
	if details.ImageURL == "" {
		t.Error("empty ImageURL")
	}
}
