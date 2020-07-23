package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := S{}
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(&s))
	fmt.Println(reflect.TypeOf(I(&s)))
	fmt.Println(reflect.TypeOf(nil))
	fmt.Println(reflect.TypeOf(I(nil)))
}

type I interface {
	F() string
}

type S struct {
}

func (s *S) F() string {
	return ""
}
