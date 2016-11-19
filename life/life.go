package life

import (
	"fmt"
	"golife/utils"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Life struct {
	config         utils.Config
	currentGenGrid [][]int
	nextGenGrid    [][]int
	gridAsString   string
	stats          utils.Stats
}

func Create(config utils.Config) Life {
	l := Life{}
	l.init(config)
	return l
}

func (life *Life) init(config utils.Config) {
	life.config = config
	life.currentGenGrid = life.create2DGrid()
	life.plantSeed()
}

func (life *Life) plantSeed() {
	for i := 0; i < life.config.NumOfSeeds; i++ {
		seed := rand.NewSource(int64(time.Now().Nanosecond()))
		randomizer := rand.New(seed)
		x := randomizer.Intn(life.config.NumOfRows)
		y := randomizer.Intn(life.config.NumOfCols)
		life.currentGenGrid[x][y] = 1
	}
}

func (life *Life) Next() {
	life.clearScreen()
	life.nextGenGrid = life.create2DGrid()
	for x := 0; x < life.config.NumOfRows; x++ {
		life.gridAsString += "\n"
		for y := 0; y < life.config.NumOfCols; y++ {
			nextCellState, currentCellState := life.getUpdatedCellState(x, y)
			life.nextGenGrid[x][y] = nextCellState
			life.gridAsString += life.createCell(currentCellState, nextCellState)
		}
	}
	life.stats.Generation++
}

func (life *Life) create2DGrid() [][]int {
	grid := make([][]int, life.config.NumOfRows)
	for x := 0; x < life.config.NumOfRows; x++ {
		grid[x] = make([]int, life.config.NumOfCols)
	}

	return grid
}

func (life *Life) getUpdatedCellState(x, y int) (int, int) {
	nextCellState := 0
	currentCellState := life.currentGenGrid[x][y]
	numOfNeighbours := life.countNeighbours(x, y)
	if (numOfNeighbours == 2 && currentCellState == 1) || numOfNeighbours == 3 {
		nextCellState = 1
	}

	return nextCellState, currentCellState
}

func (life *Life) countNeighbours(x int, y int) int {
	numOfNeighbours := 0
	startX := x - 1
	startY := y - 1

	for r := startX; r <= startX+2; r++ {
		for c := startY; c <= startY+2; c++ {
			if r >= 0 && c >= 0 &&
				r < life.config.NumOfRows &&
				c < life.config.NumOfCols &&
				(r != x || c != y) {
				numOfNeighbours += life.currentGenGrid[r][c]
			}
		}
	}

	return numOfNeighbours
}

func (life *Life) Print() {
	fmt.Println(life.gridAsString)
	fmt.Printf("\n\n\nGeneration: %d   Born: %d   Died: %d", life.stats.Generation, life.stats.Born, life.stats.Died)
	life.currentGenGrid = life.nextGenGrid
	life.gridAsString = ""
}

func (life *Life) createCell(currentCellState, nextCellState int) string {
	cellSymbol := "*"
	if currentCellState == nextCellState && nextCellState == 0 {
		cellSymbol = " "
	} else if currentCellState == nextCellState && nextCellState == 1 {
		cellSymbol = utils.ColorString(cellSymbol, "green")
	} else if currentCellState < nextCellState {
		life.stats.Born++
		cellSymbol = utils.ColorString(cellSymbol, "cyan")
	} else if currentCellState > nextCellState {
		life.stats.Died++
		cellSymbol = utils.ColorString(cellSymbol, "red")
	}
	return cellSymbol
}

func (life Life) clearScreen() {
	theCMD := exec.Command("clear")
	theCMD.Stdout = os.Stdout
	theCMD.Run()
}
