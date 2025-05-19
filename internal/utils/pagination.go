package utils

import "strconv"

const DEFAULT_PAGE_SIZE = 20

func ParsePage(pageStr string) int {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 1
	}
	return page
}
