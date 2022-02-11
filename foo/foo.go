package foo

import (
	"fmt"
	"strconv"
)

func say(n int) string {
	switch {
	case n%15 == 0:
		return "FooBar"
	case n%5 == 0:
		return "Bar"
	case n%3 == 0:
		return "Foo"
	default:
		return strconv.Itoa(n)
	}
}

func SayAny(i interface{}) string {
	if n, ok := i.(int); ok {
		return say(n)
	}
	if s, ok := i.(string); ok {
		n, _ := strconv.Atoi(s)
		return say(n)
	}
	return ""
}

type Intner interface {
	Intn(n int) int
}

func RandomSay(r Intner) string {
	return fmt.Sprintf("%s-%s-%s-%s", say(r.Intn(9)+1), say(r.Intn(9)+1), say(r.Intn(9)+1), say(r.Intn(9)+1))
}
