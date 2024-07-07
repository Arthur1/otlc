package testutil

import (
	"testing"
	"time"
)

func Time(t *testing.T, s string) time.Time {
	ti, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatal(err)
	}
	return ti
}
