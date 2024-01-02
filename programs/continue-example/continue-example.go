package main

import (
	"fmt"
	"strings"
)

func main() {
	list1 := []string{"one","two","three"}
	list2 := []string{"four","five","six"}

	for _, x := range list1 {
		if strings.EqualFold(x, "two") {
			continue
		}
		
		for _, y := range list2 {
			if strings.EqualFold(y, "five") {
				break
			}
			fmt.Printf("%s \t %s \n", x, y)
		}
	}
}