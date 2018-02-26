package ui

import "strconv"

// px is the pixel string for i.
func px(i int) string {
	if i < 0 {
		return "auto"
	}
	return strconv.Itoa(i) + "px"
}
