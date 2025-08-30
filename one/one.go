package one

import "fmt"

/* Key ideas:
* - separate your domain from "side effects"
*
* example:
* domain: string send
* side effect: fmt.Println
 */

func hello() string {
	return "some hello there"
}

func hello_args(recipient string) string {
	return "some hello there " + recipient
}

func greeting() {
	fmt.Println(hello())
}
