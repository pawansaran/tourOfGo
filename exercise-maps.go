package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for i := range words{
		_, ok := m[words[i]]
		if ok == false {
			m[words[i]] = 1
		} else {
			m[words[i]]++
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
