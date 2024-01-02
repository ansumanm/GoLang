package main

import "fmt"

type Set map[interface{}]bool

func (s Set) Add(element interface{}) {
    s[element] = true
}

func (s Set) Remove(element interface{}) {
    delete(s, element)
}

func (s Set) Contains(element interface{}) bool {
    return s[element]
}

func main() {
    mySet := make(Set)

    // Add elements to the set
    mySet.Add("apple")
    mySet.Add(42)
    mySet.Add(3.14)

    // Check if elements are in the set
    fmt.Println("Contains 'apple':", mySet.Contains("apple"))
    fmt.Println("Contains 42:", mySet.Contains(42))
    fmt.Println("Contains 2.71:", mySet.Contains(2.71))

    // Remove an element from the set
    mySet.Remove("apple")

    // Check if the removed element is still in the set
    fmt.Println("Contains 'apple':", mySet.Contains("apple"))
}

