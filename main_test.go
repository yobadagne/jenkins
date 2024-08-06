package main

import "testing"

func TestMain(t *testing.T) {
    expected := "Hello, Jenkins!"
    if got := getGreeting(); got != expected {
        t.Errorf("Expected %s, but got %s", expected, got)
    }
}

func getGreeting() string {
    return "Hello, Jenkins!"
}
