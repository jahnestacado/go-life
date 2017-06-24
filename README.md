# golife
Implementation of Conway's Game of Life in golang

The cells can be in one of three different states.
* Newborn (Cyan)
* Alive (Green)
* Dead (Red)

## Run
```shell
go run main/main.go --rows=50 --cols=150 --seeds=700
```
### Command line arguments
* rows - The number of rows in the grid [optional]
* cols - The number of columns in the grid [optional]
* seeds - The number of the initial randomly placed cells in the grid [optional]
