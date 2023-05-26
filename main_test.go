package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"divide by 1", 100, 1, 100, false},
	{"divide by 0", 100, 0, 0, true},
	{"divide by -1", 100, -1, -100, false},
}

func TestDivide(t *testing.T) {

	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected to get an error")
			}
		} else {
			if err != nil {
				t.Error("Did not expect to get an error")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f, got %f", tt.expected, got)
		}
	}
}
