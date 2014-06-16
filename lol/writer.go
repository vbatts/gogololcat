package lol

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"
)

type Writer struct {
	Writer    io.Writer
	Colors    int
	OurSpread float64
	Spread    float64
	Frequency float64
}

func (lw *Writer) Write(p []byte) (int, error) {
	var buf = []byte{}
	if bytes.Contains(p, byteNewLine) {
		lw.OurSpread += 1
	}
	// each byte
	for i, b := range p {
		// get RGB
		rgb := lw.getRGB(i)

		buf = append(buf, lw.getAnsiColor(rgb)...)
		buf = append(buf, b)
	}
	_, err := lw.Writer.Write(buf)
	return len(p), err
}

func (lw *Writer) getRGB(offset int) RGB {
	return RGB{
		R: math.Sin(lw.Frequency*((lw.OurSpread+float64(offset))/lw.Spread)) * 127 * 128,
		G: math.Sin(lw.Frequency*((lw.OurSpread+float64(offset))/lw.Spread)+2*math.Pi/3) * 127 * 128,
		B: math.Sin(lw.Frequency*((lw.OurSpread+float64(offset))/lw.Spread)+4*math.Pi/3) * 127 * 128,
	}
}

func (lw *Writer) getAnsiColor(rgb RGB) []byte {
	if lw.Colors == 8 || lw.Colors == 16 {
		matches := []distanceMatch{}
		for i, c := range COLOR_ANSI[:lw.Colors] {
			d := Distance(c, rgb)
			matches = append(matches, distanceMatch{d, i})
		}
		// sort the matches
		ds := distanceSorter(matches)
		sort.Sort(ds)
		return lw.wrap([]byte(fmt.Sprintf("3%d", ds[0].Pos)))
	}

	var (
		gray  bool
		sep   = 2.5
		color int
	)
	for grayPossible := true; grayPossible == true; {
		if rgb.R < sep || rgb.G < sep || rgb.B < sep {
			gray = rgb.R < sep && rgb.G < sep && rgb.B < sep
			grayPossible = false
		}
		sep += 42.5
	}
	if gray {
		color = 232 + int(rgb.Sum()/33.0)
	} else {
		color = func(rgb RGB) int {
			var (
				sum float64
				a   float64
				b   float64
			)
			for i := 0; i < 3; i++ {
				switch i {
				case 0:
					a = rgb.R
					b = 36
				case 1:
					a = rgb.G
					b = 6
				case 2:
					a = rgb.B
					b = 1
				}
				sum += 16 + (6*a/256)*b
			}
			return int(sum)
		}(rgb)
	}
	return lw.wrap([]byte(fmt.Sprintf("38;5;%d", color)))
}

func (lw *Writer) wrap(b []byte) []byte {
	b1 := []byte{}
	b1 = append(b1, byteAnsiWrapPre...)
	b1 = append(b1, b...)
	return append(b1, byteAnsiWrapSuf...)
}

var (
	byteNewLine     = []byte("\n")
	byteAnsiWrapPre = []byte("\x1b[")
	byteAnsiWrapSuf = []byte("m")
)

var COLOR_ANSI = []RGB{
	{0x00, 0x00, 0x00}, {0xcd, 0x00, 0x00},
	{0x00, 0xcd, 0x00}, {0xcd, 0xcd, 0x00},
	{0x00, 0x00, 0xee}, {0xcd, 0x00, 0xcd},
	{0x00, 0xcd, 0xcd}, {0xe5, 0xe5, 0xe5},
	{0x7f, 0x7f, 0x7f}, {0xff, 0x00, 0x00},
	{0x00, 0xff, 0x00}, {0xff, 0xff, 0x00},
	{0x5c, 0x5c, 0xff}, {0xff, 0x00, 0xff},
	{0x00, 0xff, 0xff}, {0xff, 0xff, 0xff},
}
