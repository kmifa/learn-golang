package strings

import (
	"slices"
	"strings"
	"testing"
	"unicode"
)

func TestToUpper(t *testing.T) {
	t.Run("upper case", func(t *testing.T) {
		got := strings.ToUpper("hello, world!")
		want := "HELLO, WORLD!"

		assertGotWant(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := strings.ToUpper("")
		want := ""

		assertGotWant(t, got, want)
	})
}

// トルコ語やアゼルバイジャン語などの “i” の変換ルールが異なる言語で重要。普段は使わない
func TestToUpperSpecial(t *testing.T) {
	t.Run("upper case 'i'", func(t *testing.T) {
		got := strings.ToUpperSpecial(unicode.TurkishCase, "i")
		want := "İ"

		assertGotWant(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := strings.ToUpperSpecial(unicode.TurkishCase, "")
		want := ""

		assertGotWant(t, got, want)
	})
}

func TestTrimPrefix(t *testing.T) {
	t.Run("prefix found", func(t *testing.T) {
		got := strings.TrimPrefix("Hello, world!", "Hello, ")
		want := "world!"

		assertGotWant(t, got, want)
	})

	t.Run("empty prefix", func(t *testing.T) {
		got := strings.TrimPrefix("Hello, world!", "")
		want := "Hello, world!"

		assertGotWant(t, got, want)
	})

	t.Run("prefix space", func(t *testing.T) {
		got := strings.TrimPrefix("  Hello, world!", "Hello")
		want := "  Hello, world!"

		assertGotWant(t, got, want)
	})
}

func TestFields(t *testing.T) {
	t.Run("fields found", func(t *testing.T) {
		got := strings.Fields("Hello, world!")
		want := []string{"Hello,", "world!"}

		if !slices.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("fields found spaces ", func(t *testing.T) {
		got := strings.Fields("Hello,  world!   Taro     ")
		want := []string{"Hello,", "world!", "Taro"}

		if !slices.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		got := strings.Fields("")
		want := []string{}

		if !slices.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestHasPrefix(t *testing.T) {
	t.Run("prefix found", func(t *testing.T) {
		got := strings.HasPrefix("Hello, world!", "Hello")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("prefix not found", func(t *testing.T) {
		got := strings.HasPrefix("Hello, world!", "world")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("empty prefix", func(t *testing.T) {
		got := strings.HasPrefix("Hello, world!", "")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func assertGotWant(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
