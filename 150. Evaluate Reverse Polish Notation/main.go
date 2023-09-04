package main

import "strconv"

// "strconv"
// "strings"

func evalRPN(tokens []string) int {
	tmp := []int{}

	for _, i := range tokens {
		switch i {
		case "+":
			tmp[len(tmp)-2] = tmp[len(tmp)-2] + tmp[len(tmp)-1]
		case "-":
			tmp[len(tmp)-2] = tmp[len(tmp)-2] - tmp[len(tmp)-1]
		case "*":
			tmp[len(tmp)-2] = tmp[len(tmp)-2] * tmp[len(tmp)-1]
		case "/":
			tmp[len(tmp)-2] = tmp[len(tmp)-2] / tmp[len(tmp)-1]
		default:
			v, _ := strconv.Atoi(i)
			tmp = append(tmp, v)
			continue
		}
		tmp = tmp[0 : len(tmp)-1]
	}
	return tmp[len(tmp)-1]
}

func main() {
	s := []string{"2", "1", "+", "3", "*"}

	evalRPN(s)
}
