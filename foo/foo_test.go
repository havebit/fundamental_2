package foo

import (
	"fmt"
	"testing"
)

func TestFooBar(t *testing.T) {
	givenWant := map[int]string{
		1:  "1",
		2:  "2",
		4:  "4",
		7:  "7",
		3:  "Foo",
		6:  "Foo",
		9:  "Foo",
		5:  "Bar",
		10: "Bar",
		20: "Bar",
		15: "FooBar",
		30: "FooBar",
		45: "FooBar",
	}

	for given, want := range givenWant {
		t.Run("", func(t *testing.T) {
			get := say(given)

			if want != get {
				t.Errorf("given %d wants %q but got %q\n", given, want, get)
			}
		})
	}
}

type fakeIntn chan int

func (f fakeIntn) Intn(n int) int {
	return <-f
}

func TestRandomSay(t *testing.T) {

	t.Run(fmt.Sprintf("given random 0,1,2,4 to RandomSay"), func(t *testing.T) {
		var random fakeIntn = make(fakeIntn, 4)
		random <- 0
		random <- 1
		random <- 2
		random <- 4
		close(random)

		want := "1-2-Foo-Bar"

		get := RandomSay(random)

		if want != get {
			t.Errorf(" wants %q but got %q\n", want, get)
		}
	})
}

// type fakeIntn chan int

// func (f fakeIntn) Intn(n int) int {
// 	return <-f
// }

// var given fakeIntn = make(chan int, 4)
// 		given <- 0
// 		given <- 1
// 		given <- 2
// 		given <- 4
// 		close(given)
