package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// go run . data.txt
	if len(os.Args[1:]) > 1 || len(os.Args[1:]) < 1 || os.Args[1] != "data.txt" {
		fmt.Fprint(os.Stderr, "Invalid input !!")
		os.Exit(1)
	}
	file := os.Args[1]
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprint(os.Stderr, "Problem with reading the file !!")
		os.Exit(1)
	}
	sli := strings.Split(strings.TrimSpace(string(data)), "\n")
	var temp []int
	for _, str := range sli {
		if str != "" {
			temp1, _ := strconv.Atoi(str)
			temp = append(temp, temp1)
		}
	}
	Linear_Regression_Line(temp)
}

func Linear_Regression_Line(Numbers []int) {
	var as, bs int
	for i := 0; i < len(Numbers)-1; i++ {
		var a int
		for j := i + 1; j < len(Numbers); j++ {
			a = (Numbers[j]-Numbers[i])/j - i

			as += a

		}
		b := Numbers[i] - (a * i)
		bs += b

	}
	fmt.Println(as, bs)
}
