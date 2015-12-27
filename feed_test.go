package main

import (
	"testing"
)

func TestFeedURLOf(t *testing.T) {
	expected := "https://coreos.com/dist/aws/aws-stable.json"
	actual := FeedURLOf("stable")

	if expected != actual {
		t.Errorf("expected: %s\ngot: %s\n", expected, actual)
	}
}
