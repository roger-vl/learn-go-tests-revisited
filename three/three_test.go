package three

import "testing"

// Iteration

/* Key ideas:
*
* Requirement: function that repeats a character 5 times
*
* New requirement: adding benchmark to improve performance
*
* New requirement: improve performance by using string builder
*
 */

func TestRepeater(t *testing.T) {
	got := Repeater("d")
	want := "ddddd"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

// to actually run this: go test -bench=. ./three/

func BenchmarkRepeated(b *testing.B) {
	// setup of test
	for b.Loop() {
		// running of code to measure
		Repeater("d")
	}
	// cleanup
}
