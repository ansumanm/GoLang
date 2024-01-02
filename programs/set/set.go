package main

import "fmt"

type Set map[string]bool

func (s Set) Add(element string) {
    s[element] = true
}

func (s Set) Remove(element string) {
    delete(s, element)
}

func (s Set) Contains(element string) bool {
    return s[element]
}

func main() {
    mySet := make(Set)

    // Add elements to the set
    mySet.Add("apple")
    mySet.Add("banana")
    mySet.Add("orange")

    // Check if an element is in the set
    fmt.Println("Contains 'apple':", mySet.Contains("apple"))
    fmt.Println("Contains 'pear':", mySet.Contains("pear"))

    // Remove an element from the set
    mySet.Remove("banana")

    // Check if the removed element is still in the set
    // fmt.Println("Contains 'banana':", mySet.Contains("banana"))
    fmt.Println("Contains 'banana':", mySet["banana"])
    fmt.Println("Contains 'papaya':", mySet["papaya"])


}
