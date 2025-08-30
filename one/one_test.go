package one

import "testing"

/* Key ideas:
*
* - go doc builtin
* example: go doc testing
* package testing // import "testing" ...
 */

func TestGreeting(t *testing.T) {
	got := hello()
	want := "some hello there"

	// learn more about placeholders, interesting
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TDD begins
// New requirement: greeting has to have a recipient
//
// Attempt:
// 1.- to many tries
// 2.- change to accept the argument but we not used already
// 3.- use the argument the test pass

func TestGreetingRecipient(t *testing.T) {
	got := hello_args("buddy")
	want := "some hello there buddy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

//-- commit after the working code
