package iteration

import (
	"fmt"
	"testing"
)

func TestREpeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q bug got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Println(result)
	// Output: aaaaaa
}
