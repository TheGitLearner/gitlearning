package strutil

import "strings"

func WordCount(input string) int{
	return len(strings.Fields(input))
}
