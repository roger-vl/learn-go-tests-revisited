package three

import "strings"

/* Key ideas:
*
* - if we now that a variable is not initializing just use var
*
* Benchmark:
*
* Use benchmark along the testing
*
* check the performance to know if an implementation is actually good
*
 */

func Repeater(digit string) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString(digit)
	}
	return repeated.String()
}
