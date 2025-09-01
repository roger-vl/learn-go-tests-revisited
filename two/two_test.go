package two

import "testing"

/* Key ideas:
*
* Requirement: add two numbers
*
*
*
*
 */

func TestAdd(t *testing.T) {
	result := AddSomething(1, 5)
	expected := 6

	if result != expected {
		t.Errorf("expected '%d' but result '%d'", expected, result)
	}
}
