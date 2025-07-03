package other

import "strings"

func CountMarketers(titles []string) int {
	count := 0
	for _, title := range titles {
		if strings.EqualFold(title, "MARKETER") {
			count++
		}
	}
	return count
}
