package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/vbatts/gogololcat/lol"
)

func main() {
	flag.Parse()

	// make sure we clear the screen :-)
	defer lol.AnsiClear(os.Stdout)

	lw := &lol.Writer{Writer: os.Stdout, OurSpread: rand.Float64(), Frequency: *flFreq, Spread: *flSpread, Colors: *flMode}

	if flag.NArg() == 0 {
		_, err := io.Copy(lw, os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	for _, arg := range flag.Args() {
		fh, err := os.Open(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = io.Copy(lw, fh)
		fh.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	flSpread = flag.Float64("s", 3.0, "Gradient spread")
	flFreq   = flag.Float64("f", 0.1, "Gradient frequency")
	flMode   = flag.Int("m", lol.DetectTermColor(), "Colors (8,16,256)")
)
