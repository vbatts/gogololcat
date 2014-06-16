package lol

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ASCIIClear is the ASCII needed to reset the color settings
var ASCIIClear = "\x1b[0m"

// AnsiClear writes the needed reset for the color settings
func AnsiClear(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s", ASCIIClear)
	return err
}

// DetectTermColor is the "Poor-mans color mode detection."
func DetectTermColor() int {
	if len(os.Getenv("ANSICON")) > 0 {
		return 16
	}
	if os.Getenv("ConEmuANSI") == "ON" {
		return 256
	}
	term := "xterm-256color"
	if len(os.Getenv("TERM")) > 0 {
		term = os.Getenv("TERM")
	}
	if strings.HasSuffix(term, "-256color") || term == "xterm" || term == "screen" {
		return 256
	}
	if strings.HasSuffix(term, "-color") || term == "rxvt" {
		return 16
	}
	return 256
}
