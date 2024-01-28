package ReverseArray

import "testing"

func TestSayHello(t *testing.T) {
	expected := "Hello, Go Guru!"
	if result := ReverseArray("Go Guru"); result != expected {
		t.Errorf("SayHello() = %q, want %q", result, expected)
	}
}
