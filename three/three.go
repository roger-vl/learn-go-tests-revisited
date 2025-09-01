package three

/* Key ideas:
*
* - if we now that a variable is not initializing just use var
*
*
 */

func Repeater(digit string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += digit
	}
	return repeated
}
