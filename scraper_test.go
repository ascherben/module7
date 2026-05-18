package main

import (
	"encoding/json"
	"testing"
)

func TestClean(t *testing.T) {
	input := " Robotics   is useful. [1] [22] "
	expected := "Robotics is useful."

	result := clean(input)

	if result != expected {
		t.Errorf("clean result: %q, but we expected: %q", result, expected)
	}
}

func TestPage(t *testing.T) {
	page := Page{
		URL:     "https://en.wikipedia.org/wiki/Robotics",
		Content: "Robotics is the study of robots.",
		Time:    "2026-01-01T00:00:00Z",
	}

	data, err := json.Marshal(page)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result["url"] != page.URL {
		t.Error("JSON should include url field")
	}

	if result["text"] != page.Content {
		t.Error("JSON should include text field")
	}

	if result["timestamp"] != page.Time {
		t.Error("JSON should include timestamp field")
	}
}
