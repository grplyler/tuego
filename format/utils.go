package format

import (
	"math"
	"strings"
)

func Fillr(data string, width int, fill string, before string, after string) string {

	dlen := len(data)
	flen := width - dlen
	filler := strings.Repeat(fill, flen)
	return before + data + filler + after

}

func Filll(data string, width int, fill string, before string, after string) string {

	dlen := len(data)
	flen := width - dlen
	filler := strings.Repeat(fill, flen)
	return before + filler + data + after

}

func ProgressBar(percent float64, width int, fill string, before string, after string) string {

	filled := percent / float64(100) * float64(width)
	// filled = int(filled)
	empty := width - int(filled)
	rfilled := int(math.Round(filled))
	return before + strings.Repeat(fill, rfilled) + strings.Repeat(" ", empty) + after
}

func PrintClassHeader() {

}
