package context

import "testing"

func TestB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"t1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			B()
		})
	}
}
