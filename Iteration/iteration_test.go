package main

import "testing"

func TestRepeat(t *testing.T)  {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatWithBuilder(b *testing.B) {
	for b.Loop() {
		RepeatWithBuilder("a", 5)
	}
}