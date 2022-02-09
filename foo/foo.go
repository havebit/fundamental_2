package foo

import "strconv"

func say(n int) string {
	if n == 6 {
		return "Foo"
	}
	if n == 5 {
		return "Bar"
	}

	if n == 3 {
		return "Foo"
	}
	return strconv.Itoa(n)
}
