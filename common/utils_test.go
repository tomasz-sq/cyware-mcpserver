package common

import (
	"encoding/json"
	"math"
	"testing"
	"time"
)

func TestParseTimeoutDurationAcceptsNumericRepresentations(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want time.Duration
	}{
		{name: "float64", arg: float64(1.5), want: 1500 * time.Millisecond},
		{name: "float32", arg: float32(2), want: 2 * time.Second},
		{name: "int", arg: int(3), want: 3 * time.Second},
		{name: "int64", arg: int64(4), want: 4 * time.Second},
		{name: "uint64", arg: uint64(5), want: 5 * time.Second},
		{name: "json number", arg: json.Number("6.25"), want: 6250 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTimeoutDuration(tt.arg)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestParseTimeoutDurationRejectsInvalidValues(t *testing.T) {
	tests := []struct {
		name string
		arg  any
	}{
		{name: "string", arg: "1"},
		{name: "bad json number", arg: json.Number("bad")},
		{name: "zero", arg: 0},
		{name: "negative", arg: -1},
		{name: "nan", arg: math.NaN()},
		{name: "infinity", arg: math.Inf(1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseTimeoutDuration(tt.arg); err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}
