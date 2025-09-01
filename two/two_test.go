package two

import (
	"fmt"
	"testing"
)

/* Key ideas:
*
* - We can have examples as part of the testing files
* automatically added to docs and running along the actual tests
*
* Requirement: add two numbers
* New requirement: add an example
*
 */

func TestAdd(t *testing.T) {
	result := AddSomething(1, 5)
	expected := 6

	if result != expected {
		t.Errorf("expected '%d' but result '%d'", expected, result)
	}
}

func ExampleAddSomething() {
	result := AddSomething(1, 5)
	fmt.Println(result)
	// output: 6
}
