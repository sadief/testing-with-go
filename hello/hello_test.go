package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lassie")
	want := "Hello, Lassie"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
