package main

import "fmt"

type iInterface interface {
	aFunc()
}

type sStruct struct {
}

func (s sStruct)aFunc() {
	fmt.Println("Inside s")
}

func NewFunc(i iInterface) {
	i.aFunc()
}

func AFunc() {
	var tmp sStruct
	NewFunc(tmp)
}


