package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortstring(dict map[string]int) []string {
	var s []string
	for k, val := range dict {
		a := fmt.Sprint(k, val)
		s = append(s, a)
	}
	sort.Strings(s)
	return s
}

func group(str string) string {
	arr := strings.Split(str, "")
	dict := make(map[string]int)
	for _, num := range arr {
		i := string(num)
		dict[i] = dict[i] + 1
	}

	s := sortstring(dict)
	res := strings.Join(s, "")

	return res
}

func main() {
	a := "aaafffbbbpppppAAAAAkkkkkBBBjjjjjjjjjjjTTT"
	res := group(a)
	fmt.Print(res)
}
