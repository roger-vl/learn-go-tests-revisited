package three

import "testing"

// Iteration

/* Key ideas:
*
* Requirement: function that repeats a character 5 times
*
 */

func TestRepeater(t *testing.T) {
	got := Repeater("d")
	want := "ddddd"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
