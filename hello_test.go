package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jesus")
	want := "Hello, Jesus!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
