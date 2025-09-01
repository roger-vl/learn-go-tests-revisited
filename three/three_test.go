package three

import (
	"fmt"
	"testing"
)

// Iteration

/* Key ideas:
*
* Requirement: function that repeats a character 5 times
*
* New requirement: adding benchmark to improve performance
*
* New requirement: improve performance by using string builder
*
* New requirement: the caller can specify how many time the digit is repeated
* write an example
*
 */

func TestRepeater(t *testing.T) {
	t.Run("repeated 5 times d", func(t *testing.T) {
		got := Repeater("d", 5)
		want := "ddddd"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("repeated 10 times d", func(t *testing.T) {
		got := Repeater("d", 10)
		want := "dddddddddd"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

// to actually run this: go test -bench=. ./three/
// the use of "-benchmem" gives stats about memory allocations

func BenchmarkRepeated(b *testing.B) {
	// setup of test
	for b.Loop() {
		// running of code to measure
		Repeater("d", 5)
	}
	// cleanup
}

// example
func ExampleRepeater() {
	repeated := Repeater("a", 4)
	fmt.Println(repeated)
	// output: aaaa
}
