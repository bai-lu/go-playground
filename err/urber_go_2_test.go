// package foo
package err

import "testing"

func TestF2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"F2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F2()
		})
	}
}
