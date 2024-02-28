package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mal")
	want := "Hello, Mal"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
