package expenses

import "fmt"

func fib() func() int {
	var n1, n2 int

	return func() int {
		if n1 == 0 && n2 == 0 {
			n1 = 1
		} else {
			n1, n2 = n2, n1+n2
		}
		return n2
	}
}

func main() {
	next := fib()
	for i := 0; i < 9; i++ {
		fmt.Printf("F%d\t= %4d\n", i, next())
	}
}

/**
F0	=    0
F1	=    1
F2	=    1
F3	=    2
F4	=    3
F5	=    5
F6	=    8
F7	=   13
F8	=   21
*/
