package main

import (
	"flag"
	"image/color"
	"time"

	"go-life/life"
	"go-life/utils"

	"github.com/mcuadros/go-rpi-rgb-led-matrix"
)

const interval = 1000

var (
	numRowsPtr  = flag.Int("rows", 32, "The number of rows")
	numColsPtr  = flag.Int("cols", 64, "The number of cols")
	numSeedsPtr = flag.Int("seeds", 200, "The number of initial random seeds")
)

func main() {
	flag.Parse()
	config := utils.Config{
		NumOfRows:  *numRowsPtr,
		NumOfCols:  *numColsPtr,
		NumOfSeeds: *numSeedsPtr,
	}
	l := life.New(config)

	matrixConfig := &rgbmatrix.DefaultConfig
	matrixConfig.Rows = *numRowsPtr
	matrixConfig.Cols = *numColsPtr
	matrixConfig.Parallel = 1
	matrixConfig.ChainLength = 1
	matrixConfig.Brightness = 100
	matrixConfig.HardwareMapping = "regular"
	matrixConfig.ShowRefreshRate = false
	matrixConfig.InverseColors = false
	matrixConfig.DisableHardwarePulsing = false

	m, err := rgbmatrix.NewRGBLedMatrix(matrixConfig)
	if err != nil {
		panic(err)
	}

	c := rgbmatrix.NewCanvas(m)
	defer c.Close()

	for {
		l.Next()
		grid := l.GetNextGenGrid()

		for x := 0; x < *numRowsPtr; x++ {
			for y := 0; y < *numColsPtr; y++ {
				cell := grid[x][y]

				cellColor := color.RGBA{0, 0, 0, 255}
				if cell.Color == "green" {
					cellColor = color.RGBA{0, 255, 0, 255}
				}
				if cell.Color == "cyan" {
					cellColor = color.RGBA{25, 255, 255, 255}
				}
				if cell.Color == "red" {
					cellColor = color.RGBA{255, 0, 0, 255}
				}
				c.Set(y, x, cellColor)
			}
		}

		c.Render()

		time.Sleep(time.Millisecond * interval)
	}
}
