package main

import (
	"strings"
	"testing"
	"time"
)

var (
	tests = []struct {
		title    string
		log      string
		count    int
		duration time.Duration
	}{
		{
			"one interval",
			`2015-03-06 16:13:00.000 +00:00 [Information] Starting A.
2015-03-06 16:13:20.000 +00:00 [Information] Finished A.`,
			1,
			time.Duration(20) * time.Second,
		},
		{
			"two intervals",
			`2015-03-06 16:13:00.000 +00:00 [Information] Starting A.
2015-03-06 16:13:20.000 +00:00 [Information] Finished A.
2015-03-06 16:14:00.000 +00:00 [Information] Starting A.
2015-03-06 16:14:40.000 +00:00 [Information] Finished A.`,
			2,
			time.Duration(30) * time.Second,
		},
		{
			"wrong intervals are ignored",
			`2015-03-06 16:12:00.000 +00:00 [Information] Starting A.
2015-03-06 16:12:20.000 +00:00 [Information] Finished A.
2015-03-06 16:13:00.000 +00:00 [Information] Starting A.
2015-03-06 16:14:00.000 +00:00 [Information] Starting A.
2015-03-06 16:14:20.000 +00:00 [Information] Finished A.
2015-03-06 16:14:40.000 +00:00 [Information] Finished A.
2015-03-06 16:15:00.000 +00:00 [Information] Starting A.
2015-03-06 16:15:40.000 +00:00 [Information] Finished A.`,
			2,
			time.Duration(30) * time.Second,
		},
	}
)

func TestSnapshotTimeSimpleInterval(t *testing.T) {
	m := measure{"Starting A.", "Finished A.", "2006-01-02 15:04:05.000", 0, 23, strings.NewReader("")}
	for _, test := range tests {
		t.Log(test.title)
		m.reader = strings.NewReader(test.log)
		duration, count := m.Intervals()

		if count != test.count {
			t.Errorf("Expected: %d, got %d", test.count, count)
		}

		if duration != test.duration {
			t.Errorf("Expected: %v, got %v", test.duration, duration)
		}
	}
}
