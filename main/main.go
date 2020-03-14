package main

import (
	"flag"
	"time"

	"go-life/life"
	"go-life/utils"
)

const interval = 500

var (
	numRowsPtr  = flag.Int("rows", 40, "The number of rows")
	numColsPtr  = flag.Int("cols", 150, "The number of cols")
	numSeedsPtr = flag.Int("seeds", 500, "The number of initial random seeds")
)

func main() {
	flag.Parse()
	config := utils.Config{
		NumOfRows:  *numRowsPtr,
		NumOfCols:  *numColsPtr,
		NumOfSeeds: *numSeedsPtr,
	}
	l := life.New(config)
	for {
		time.Sleep(time.Millisecond * interval)
		l.Next()
		l.Print()
	}
}
