package main

import (
	"flag"
	"time"

	life "github.com/jahnestacado/go-life"
)

const interval = 500

var (
	numRowsPtr  = flag.Int("rows", 32, "The number of rows")
	numColsPtr  = flag.Int("cols", 64, "The number of cols")
	numSeedsPtr = flag.Int("seeds", 200, "The number of initial random seeds")
)

func main() {
	flag.Parse()
	config := life.Config{
		NumOfRows:  *numRowsPtr,
		NumOfCols:  *numColsPtr,
		NumOfSeeds: *numSeedsPtr,
	}
	l := life.New(config)

	for {
		l.Next()
		l.Print()

		time.Sleep(time.Millisecond * interval)
	}
}
