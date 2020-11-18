package colortest

import "testing"

func TestToRed(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		text string
		want string
	}{
		"hello to red hello": {
			"hello",
			"\x1b[31mhello\x1b[0m",
		},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := ToRed(tc.text)
			if got != tc.want {
				t.Errorf("%s (got: %q, want: %q)", name, got, tc.want)
			}
		})
	}
}
