package one

/* Key ideas:
* - separate your domain from "side effects"
* - use of constants to improve performance
*
* example:
* domain: string send
* side effect: fmt.Println
* The cicle of TDD:
*
* 1. Write a test
* 2. Make the compiler pass
* 3. Run the test, should fail, check the message is meaningful
* 4. Write enough code to make the test pass
* 5. Refactor
*
* > Sticking to the "feedback loop"
*
 */

const (
	spanish  = "es"
	french   = "fr"
	japanese = "ja"

	en_greet = "some hello there, "
	es_greet = "hola tu, "
	fr_greet = "Bonjour, "
	ja_greet = "Konnichiwa, "
)

func hello() string {
	return "some hello there"
}

func hello_args(recipient, lang string) string {
	if recipient == "" {
		recipient = "world"
	}
	return greetByLang(lang) + recipient
}

func greetByLang(lang string) (greet string) {
	switch lang {
	case spanish:
		greet = es_greet
	case french:
		greet = fr_greet
	case japanese:
		greet = ja_greet
	default:
		greet = en_greet
	}
	return
}
