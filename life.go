package life

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"go-life/utils"
)

const symbol = "*"

var (
	codes map[string]int
)

func init() {
	codes = map[string]int{
		"red":   31,
		"green": 32,
		"cyan":  36,
	}
}

func getColoredSymbol(color string) string {
	return "\033[" + strconv.Itoa(codes[color]) + "m" + symbol + "\033[0m"
}

type Life interface {
	Next()
	Print()
	GetGrid() [][]Cell
}

type Cell struct {
	State int
	Color string
}

type Stats struct {
	Generation int
	Born       int
	Died       int
}

type Config struct {
	NumOfRows  int
	NumOfCols  int
	NumOfSeeds int
}

type life struct {
	config         Config
	currentGenGrid [][]Cell
	nextGenGrid    [][]Cell
	gridAsString   string
	stats          Stats
}

func New(config utils.Config) Life {
	l := &life{}
	l.init(config)
	return l
}

func (life *life) init(config utils.Config) {
	life.config = config
	life.currentGenGrid = life.create2DGrid()
	life.plantSeed()
}

func (life *life) plantSeed() {
	numOfSeeds := life.config.NumOfSeeds

	if life.stats.Generation != 0 {
		numOfSeeds = int(float64(life.config.NumOfSeeds) * 0.1)
	}
	for i := 0; i < numOfSeeds; i++ {
		seed := rand.NewSource(int64(time.Now().Nanosecond()))
		randomizer := rand.New(seed)
		x := randomizer.Intn(life.config.NumOfRows)
		y := randomizer.Intn(life.config.NumOfCols)
		life.currentGenGrid[x][y] = Cell{State: 1}
	}
}

func (life *life) Next() {
	life.clearScreen()
	life.nextGenGrid = life.create2DGrid()

	fmt.Println(life.stats.Generation)
	if life.stats.Generation%100 == 0 {
		life.plantSeed()
	}

	for x := 0; x < life.config.NumOfRows; x++ {
		for y := 0; y < life.config.NumOfCols; y++ {
			nextCell := life.getNextCellState(x, y)
			life.nextGenGrid[x][y] = nextCell
		}
	}
	life.stats.Generation++

	life.currentGenGrid = life.nextGenGrid
}

func (life *life) GetGrid() [][]Cell {
	return life.nextGenGrid
}

func (life *life) create2DGrid() [][]Cell {
	grid := make([][]Cell, life.config.NumOfRows)
	for x := 0; x < life.config.NumOfRows; x++ {
		grid[x] = make([]Cell, life.config.NumOfCols)
	}

	return grid
}

func (life *life) getNextCellState(x, y int) Cell {
	nextCell := Cell{State: 0}
	currentCell := life.currentGenGrid[x][y]
	numOfNeighbours := life.countNeighbours(x, y)
	if (numOfNeighbours == 2 && currentCell.State == 1) || numOfNeighbours == 3 {
		nextCell = Cell{State: 1}
	}

	nextCell.Color = life.getColor(currentCell, nextCell)

	return nextCell
}

func (life *life) countNeighbours(x int, y int) int {
	numOfNeighbours := 0
	startX := x - 1
	startY := y - 1

	for r := startX; r <= startX+2; r++ {
		for c := startY; c <= startY+2; c++ {
			if r >= 0 && c >= 0 &&
				r < life.config.NumOfRows &&
				c < life.config.NumOfCols &&
				(r != x || c != y) {
				numOfNeighbours += life.currentGenGrid[r][c].State
			}
		}
	}

	return numOfNeighbours
}

func (life *life) Print() {
	var output string
	for x := 0; x < life.config.NumOfRows; x++ {
		output += "\n"
		for y := 0; y < life.config.NumOfCols; y++ {
			cell := life.nextGenGrid[x][y]
			symbol := " "
			if cell.Color != "black" {
				symbol = utils.GetColoredSymbol(cell.Color)
			}
			output += symbol
		}
	}
	fmt.Println(output)
	fmt.Printf("\n\n\nGeneration: %d   Born: %d   Died: %d", life.stats.Generation, life.stats.Born, life.stats.Died)
}

func (life *life) getColor(currentCell, nextCell Cell) string {
	var color string
	if currentCell.State == nextCell.State && nextCell.State == 0 {
		color = "black"
	} else if currentCell.State == nextCell.State && nextCell.State == 1 {
		color = "green"
	} else if currentCell.State < nextCell.State {
		life.stats.Born++
		color = "cyan"
	} else if currentCell.State > nextCell.State {
		life.stats.Died++
		color = "red"
	}

	return color
}

func (life life) clearScreen() {
	theCMD := exec.Command("clear")
	theCMD.Stdout = os.Stdout
	theCMD.Run()
}
