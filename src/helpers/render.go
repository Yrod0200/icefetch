package helpers

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

func maxWidth(lines []string) int {

	max := 0
	for _, l := range lines {
		w := runewidth.StringWidth(l)
		if w > max {
			max = w

		}
	}

	return max
}

func padRight(s string, width int) string {
	w := runewidth.StringWidth(s)

	if w >= width {
		return s
	}

	return s + strings.Repeat(" ", width-w)

}

func getMaxHeight(logo int, info int) int {

	if logo > info {
		return logo
	} else {
		return info
	}

}

func Render(logo []string, info []string) {
	width := maxWidth(logo)

	for i := 0; i < getMaxHeight(len(logo), len(info)); i++ {

		if len(logo) < len(info) {

			if i < len(logo) {
				left := padRight(logo[i], width)
				fmt.Println(left, info[i])
			} else {
				fmt.Println(strings.Repeat(" ", width), info[i])
			}

		} else {
			if i < len(info) {
				left := padRight(logo[i], width)
				fmt.Println(left, info[i])
			} else {
				fmt.Println(logo[i])
			}
		}

	}

}
