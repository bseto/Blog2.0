package main

import "fmt"

type Helloer interface {
	Hello[T any] (T) 
	TheThing[T any]() T
}

type A struct { }

func (a A) Hello(someString string) {
	fmt.Printf("hello: %v", someString)
}

func (a A) TheThing() string {
	return "this is my payload"
}

type B struct { }

func (b B) Hello(someInt int) {
	fmt.Printf("hello: %v", someInt)
}

func (b B) TheThing() int {
	return 5
}

func main() {

	mapper := make(map[string]Helloer)
	mapper["string"] = A{}
	mapper["int"] = B{}

	for _, helloer := range mapper {
		helloer.Hello(helloer.TheThing())
	}

	fmt.Printf("hello world")
	myStrings := []string {
		"hello",
		"world",
		"i",
		"am",
		"trying",
		"generics",
	}
	myInts := []int {1, 2, 3, 4, 5}
	myFloats := []float32 {1.0, 2.0, 3.0}
	fmt.Printf("before: %v \n",myStrings)
	Reverse(myStrings)
	fmt.Printf("after: %v \n",myStrings)

	fmt.Printf("before: %v \n",myInts)
	Reverse(myInts)
	fmt.Printf("after: %v \n",myInts)

	fmt.Printf("before: %v \n",myFloats)
	Reverse(myFloats)
	fmt.Printf("after: %v \n",myFloats)
}

func Reverse[T any](s []T) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}