package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/pallat/hello/foo"
)

// struct{}
// [0]func

func main() {
	ch := make(chan int)
	quitCh := make(chan struct{})
	go fibonacciGo(ch, quitCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	quitCh <- struct{}{}
}

func fibonacciGo(ch chan int, quitCh chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-quitCh:
			fmt.Println("quit...")
			return
		}
	}
}

func newFibonacciFunc() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

func fibonacci(n int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func mainOfChan() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

var wg sync.WaitGroup

func mainOfWaitGoroutine() {
	start := time.Now()

	wg.Add(3)
	go slow("1")
	go slow("2")
	go slow("3")

	wg.Wait()
	fmt.Println(time.Since(start))
}

func slow(s string) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	log.Println(s)
}

type IntnFunc func(int) int

func (f IntnFunc) Intn(n int) int {
	return f(n)
}

func mainOfMethodOnFunc() {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)

	var fn IntnFunc = r.Intn

	fmt.Println(foo.RandomSay(fn))

	fmt.Println(foo.RandomSay(IntnFunc(func(int) int { return 3 })))
}

func mainOfBoom() {
	fn := boom()

	fmt.Println(fn(), fn(), fn(), fn(), fn())
}

func boom() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

type addFunc func(int, int) int

func printSum(add addFunc) {
	fmt.Printf("summary: %v\n", add(64, 32))
}

func newAddFunc() addFunc {
	return add
}

func add(a, b int) int {
	return a + b
}

// func Pic(dx, dy int) [][]uint8 {
// 	return [][]uint8{}
// }

// func Show(pic func(dx, dy int) [][]uint8) {

// }

type rectangle struct {
	w, h float64
}

func (r rectangle) area() float64 {
	return r.h * r.h
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

func (s square) String() string {
	return "square"
}

type areaer interface {
	area() float64
}

func printArea(r areaer) {
	fmt.Printf("area is %v\n", r.area())
}

func mainPrintArea() {
	r := rectangle{w: 4, h: 5}
	printArea(r)

	s := square{a: 4}
	printArea(s)
}

func mainOfStringer() {
	var s String

	fmt.Println(s)
}

type error interface {
	Error() string
}

type String struct{}

func (String) Error() string {
	return "My String Type"
}

func condition(n int) {
	switch {
	case n%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}
}

type any interface{}

func mainOfAny() {
	var i any

	fmt.Println(i == nil)

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)

	if s, ok := i.(int); ok {
		fmt.Printf("type is %T, value is %v\n", s, s)
	}
	if s, ok := i.(string); ok {
		fmt.Printf("type is %T, value is %v\n", s, s)
	}
}

func mainOfDefer() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	log.Panic("hey!!!")
}

func greeting(names ...string) string {
	if len(names) == 0 {
		return "Hello"
	}
	if len(names) == 1 {
		return fmt.Sprintf("Hello %s", names[0])
	}
	if len(names) == 2 {
		return fmt.Sprintf("Hello %s and %s", names[0], names[1])
	}

	return "Hello " + strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
}

func cal(n int) {
	defer fmt.Println(n)
	defer func(i int) {
		fmt.Println(i)
	}(n)

	n += n
	fmt.Println(n)
}
