package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, world!",
		})
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

type Rectangle struct {
	Width  float64 `json:"width"`
	Length float64 `json:"length"`
}

func Area(r Rectangle) float64 {
	return r.Width * r.Length
}

func jsonMessage() {
	// rec := Rectangle{Width: 4, Length: 5}
	rec := map[string]string{
		"width":  "12",
		"length": "13",
	}

	b, err := json.Marshal(&rec)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(b))
}

const colName = 3

func oscarMorethanTwo() {
	f, err := os.Open("oscar.csv")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	nameCount := map[string]int{}

	for _, record := range records {
		nameCount[record[colName]]++
	}

	for name, count := range nameCount {
		if count > 1 {
			fmt.Println(name, count)
		}
	}

}

func mapKV() {
	m := map[string]int{}

	m["a"]++

	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
}

func vary(x ...int) {
	for _, v := range x {
		fmt.Println(v)
	}
}

func couple(s string) (r []string) {
	for s += "*"; len(s) > 1; s = s[2:] {
		r = append(r, s[:2])
	}
	return
}

func array() {
	var a = [...]int{1, 2, 3, 99}

	fmt.Printf("%T\n", a)

	for _, v := range a {
		fmt.Println(v)
	}
}

func pointer() {
	var p = new(int)

	fmt.Println(p == nil)

	i := 42
	p = &i

	fmt.Println(p, &i)
	fmt.Println(*p, i)

	*p = 43
	fmt.Println(*p, i)
}

// func power(b, x int) int {

// }

func prime(n int) {
	for i := 2; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
	}
}

func IsCorrect() bool {
	return true
}

func swap(a, b int) (int, int, bool) {
	return b, a, true
}

func squareArea(a float64) float64 {
	return a * a
}
