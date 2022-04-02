package utils

import (
	"strings"
)

// difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func Strip(input [][]Tag) string {
	output := ``

	for _, item := range input {
		for _, elem := range item {
			output = output + elem.Value + "\n"
		}
	}

	output = strings.TrimSpace(output)

	return output
}

func Rotate(input map[string][]string) [][]string {

	max := 0

	for _, value := range input {
		if len(value) > max {
			max = len(value)
		}
	}

	output := make([][]string, max+1)

	for key, value := range input {
		for n := range output {
			if n == 0 {
				output[n] = append(output[n], key)
			} else if len(value) >= n {
				output[n] = append(output[n], value[n-1])
			} else {
				output[n] = append(output[n], "")
			}
		}
	}

	return output
}
