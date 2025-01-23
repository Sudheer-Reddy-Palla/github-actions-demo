package main

import "testing"

func TestMain(t *testing.T) {
    want := "Hello, GitHub Actions!"
    got := "Hello, GitHub Actions!"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
