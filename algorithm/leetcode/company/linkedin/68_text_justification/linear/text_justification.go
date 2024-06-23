package linear

import (
	"math"
	"strings"
)

// Time: O(n)
// Space: O(n)
func fullJustify(words []string, maxWidth int) []string {
	text := []string{}
	line := []string{}
	lineCnt := 0
	textOnlyCnt := 0
	for i, word := range words {
		if lineCnt+len(word) > maxWidth {
			var str strings.Builder
			totalNumSpace := maxWidth - textOnlyCnt
			numIntervals := maxInt(len(line)-1, 1)
			space := totalNumSpace / numIntervals
			extra := totalNumSpace % numIntervals
			for _, w := range line {
				str.WriteString(w)
				for j := 0; j < space && totalNumSpace > 0; j++ {
					str.WriteString(" ")
				}
				totalNumSpace -= space
				if extra > 0 {
					str.WriteString(" ")
					totalNumSpace--
					extra--
				}
			}
			text = append(text, str.String())
			lineCnt = 0
			textOnlyCnt = 0
			line = []string{}
		}
		if word != "" {
			textOnlyCnt += len(word)
			lineCnt += len(word) + 1
			line = append(line, word)
		}
		if i == len(words)-1 {
			var str strings.Builder
			for j, w := range line {
				str.WriteString(w)
				if j != len(line)-1 {
					str.WriteString(" ")
				}
			}
			for j := 0; j < maxWidth-len(str.String()); i++ {
				str.WriteString(" ")
			}
			text = append(text, str.String())
		}
	}
	return text
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
