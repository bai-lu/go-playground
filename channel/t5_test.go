package channel

import "testing"

func Test_t5(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"t5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t5()
		})
	}
}
