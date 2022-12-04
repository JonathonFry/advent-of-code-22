package main

import "strings"

func readInputAsLines(input string) []string {
	trimmed := strings.TrimSuffix(input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	return strings.Split(trimmed, "\n")
}

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

func intersect[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func intersectThree[T comparable](a []T, b []T, c []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) && containsGeneric(c, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
