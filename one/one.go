package one

/* Key ideas:
* - separate your domain from "side effects"
* - use of constants to improve performance
*
* example:
* domain: string send
* side effect: fmt.Println
 */

const greet = "some hello there, "

func hello() string {
	return "some hello there"
}

func hello_args(recipient string) string {
	if recipient == "" {
		recipient = "world"
	}
	return greet + recipient
}
