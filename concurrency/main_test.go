package concurrency

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go printNum()
			go printChar()
			time.Sleep(1 * time.Second)
		})
	}
}
