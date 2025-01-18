package strings

import (
	"strings"
	"testing"
)

func TestTrimPrefix(t *testing.T) {
	t.Run("prefix found", func(t *testing.T) {
		got := strings.TrimPrefix("Hello, world!", "Hello, ")
		want := "world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("empty prefix", func(t *testing.T) {
		got := strings.TrimPrefix("Hello, world!", "")
		want := "Hello, world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("prefix space", func(t *testing.T) {
		got := strings.TrimPrefix("  Hello, world!", "Hello")
		want := "  Hello, world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
