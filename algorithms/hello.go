package main

import (
	"fmt"
	"strings"

	"github.com/scommerce/algorithms/io"
)

var p = fmt.Println

func main() {
	p("Solution #1: using map")

	var bigOTime int = 0
	var bigOSpace int = 0

	p("Read input file to array")
	var a string = readfile.ReadFile()
	var lenArray int = len(strings.Split(a, " "))

	Mymap := make(map[string]bool)

	p("check if value exists in map, if not, add it to map, if exists, remove it from map")
	for i := 0; i < lenArray; i++ {
		value := strings.Split(a, " ")[i]

		p("update map with", value)
		if Mymap[value] == false {
			Mymap[value] = true
		} else {
			delete(Mymap, value)
		}
		p(Mymap)

		bigOTime++
		bigOSpace++
	}
	p("")
	//p("The result: key value exists in map is the unique number")
	//p(Mymap)

	keys := make([]string, 0, len(Mymap))

	for k, _ := range Mymap {
		keys = append(keys, k)
	}

	fmt.Println("Result:", keys[0])

	p("Time:", bigOTime, "unit")
	p("Space:", bigOSpace, "unit")
}
