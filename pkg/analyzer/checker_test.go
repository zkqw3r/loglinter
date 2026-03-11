package analyzer

import (
	"testing"
)

func TestCheckUppercase(t *testing.T) {
	cases := []struct {
		input   string
		wantFix string
		wantErr bool
	}{
		{"Server started", "server started", true},
		{"server started", "server started", false},
		{"123 started", "123 started", false},
	}
	for _, c := range cases {
		r := &Result{Log: c.input}
		checkUppercase(r)
		if (len(r.Messages) > 0) != c.wantErr {
			t.Errorf("input=%q: wantErr=%v", c.input, c.wantErr)
		}
		if r.Log != c.wantFix {
			t.Errorf("input=%q: wantFix=%q got=%q", c.input, c.wantFix, r.Log)
		}
	}
}
