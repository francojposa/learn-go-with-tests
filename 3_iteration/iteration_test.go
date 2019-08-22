package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := RepeatChar("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("\nwant: %q\n got: %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("a", 5)
	}
}

func ExampleRepeat() {
	repeatedChar := RepeatChar("A", 10)
	fmt.Println(repeatedChar)
	// Output: AAAAAAAAAA
}
