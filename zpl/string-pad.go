package zpl

import "strings"

func padStart(input string, length int, padStr string) string {
	if length <= len(input) {
		return input
	}
	return strings.Repeat(padStr, length-len(input)) + input
}

// func padEnd(input string, length int, padStr string) string {
// 	if length <= len(input) {
// 		return input
// 	}
// 	return input + strings.Repeat(padStr, length-len(input))
// }
