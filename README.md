# golife

Implementation of Conway's Game of Life in golang

![Preview](https://github.com/jahnestacado/golife/blob/master/resources/golife.gif?raw=true)

The cells can be in one of three different states.

- Newborn (Cyan)
- Alive (Green)
- Dead (Red)

## Run

```shell
go run examples/main.go --rows=32 --cols=64 --seeds=300
```

### Command line arguments

- rows - The number of rows in the grid [optional]
- cols - The number of columns in the grid [optional]
- seeds - The number of the initial randomly placed cells in the grid [optional]
