package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertValue(t, repeated, expected)
	})

	t.Run("repeat b 10 times", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"
		assertValue(t, repeated, expected)
	})
}

func assertValue(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 3)
	fmt.Println(repeated)
	// Output: ccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
