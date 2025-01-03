package iteration

import (
	"fmt"
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {
	t.Run("repeat character 5 times", func(t *testing.T) {
		repeated := Repeat("a", repeatCount)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	// Test with zero times
	t.Run("repeat character 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", repeatCount)
	fmt.Println(repeated)
	// Output: aaaaa
}
