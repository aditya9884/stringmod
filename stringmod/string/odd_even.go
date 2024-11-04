// writting the statement for the

package string

func InternalFunction() {
	// In go a  neme is export (visible outside the packge)
	// if it begins with a capital letters
}

//Must begins with capital letter in order to export

func CountOddEven(s string) (odds, even int) {
	odds, even = 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			even++
		} else {
			odds++
		}
	}
	return
}
