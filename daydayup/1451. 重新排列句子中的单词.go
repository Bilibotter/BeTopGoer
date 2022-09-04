package main

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	split := strings.Split(text, " ")
	split[0] = strings.ToLower(split[0])
	sort.SliceStable(split, func(i, j int) bool {
		return len(split[i]) < len(split[j])
	})
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}
