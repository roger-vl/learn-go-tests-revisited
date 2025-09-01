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
//
//-- commit after the working code
//
// New requirement: when is called empty should say "some hello there, world"
//
// Refactor after running with default for string empty
//
// New requirement: support a second parameter, specifying the language of the greeting
// - if language is not supported default to English

func TestGreetingRecipient(t *testing.T) {
	t.Run("say hello to recipient", func(t *testing.T) {
		got := hello_args("buddy", "en")
		want := "some hello there, buddy"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say hello default to world", func(t *testing.T) {
		got := hello_args("", "en")
		want := "some hello there, world"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := hello_args("Pepe", "es")
		want := "hola tu, Pepe"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say hello in English by default", func(t *testing.T) {
		got := hello_args("Pepe", "gr")
		want := "some hello there, Pepe"
		assertCorrectMsg(t, got, want)
	})
}

// TB is an interface for T an B, witch are normal tests and benchmarks
// When 2 arguments are of the same type the declaration can be shorten to (arg1, arg2 type)
func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
