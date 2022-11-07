package test

import (
	"study-go/src"
	"testing"
)

func TestHello(t *testing.T) {
	got := src.Hello("Hello, Chris")
	want := "say:Hello, Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func TestHello2(t *testing.T) {

	funcName := func(t *testing.T, got string, want string) {
		//t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := src.Hello("")
		want := "say:Hello, Chris"

		funcName(t, got, want)
	})

	t.Run("say hello world when an Hello, Chris is supplied", func(t *testing.T) {
		got := src.Hello("Hello, Chris")
		want := "say:Hello, Chris"

		funcName(t, got, want)
	})
}
