package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestFields(t *testing.T) {
	result := strings.Fields("i    am bob.        ")
	expected := []string{"i", "am", "bob."}

	if len(result) != len(expected) {
		t.Errorf("expected %q but got %q", expected, result)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("expected %q but got %q", expected, result)
		}
	}
}

func ExampleFields() {
	result := strings.Fields("i    am bob.")
	fmt.Printf("%q", result)
	// Output: ["i" "am" "bob."]
}
