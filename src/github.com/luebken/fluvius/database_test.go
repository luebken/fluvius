package main

import (
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {
	if len(db.AllItems()) != 0 {
		t.Fail()
	}
	db.save <- Item{"test", "test", "test", "test", "test"}
	<-time.After(time.Duration(1 * time.Second))
	if len(db.AllItems()) != 1 {
		t.Fail()
	}
}
