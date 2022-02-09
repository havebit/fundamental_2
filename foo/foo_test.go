package foo

import "testing"

func TestFooBarGivenInt1WantsString1(t *testing.T) {
	given := 1
	wants := "1"

	get := say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGivenInt2WantsString2(t *testing.T) {
	given := 2
	wants := "2"

	get := say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGivenInt3WantsStringFoo(t *testing.T) {
	given := 3
	wants := "Foo"

	get := say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}
