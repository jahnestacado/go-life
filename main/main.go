package main

import (
	"flag"
	"golife/life"
	"golife/utils"
	"time"
)

const INTERVAL = 500

var NUM_ROWS_PTR = flag.Int("rows", 40, "The number of rows")
var NUM_COLS_PTR = flag.Int("cols", 150, "The number of cols")
var NUM_OF_SEEDS_PTR = flag.Int("seeds", 500, "The number of initial random seeds")

func main() {
	flag.Parse()
	config := utils.Config{*NUM_ROWS_PTR, *NUM_COLS_PTR, *NUM_OF_SEEDS_PTR}
	l := life.Create(config)
	for {
		time.Sleep(time.Millisecond * INTERVAL)
		l.Next()
		l.Print()
	}
}
