package readfile

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	strings_data := strings.TrimSuffix(string(data), "\n")
	return strings_data
}
