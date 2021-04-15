# Tile Placement CSP Backtracking

AI project. Landscape, the number of tile types ("OUTER_BOUNDARY", "EL_SHAPE", "FULL_BLOCK"), and targets are given. The program returns the correct tile placement on the landscape to meet the target.

## Installation

Go should be installed beforehand https://golang.org/doc/install and GOPATH should be set.

Then either download .zip file of this package or go get from command line to GOPATH:

```
$ go get github.com/shahaliyev/go-csp-tp
```

### Usage

If the installation part is complete, it remains to run the command line with the path to the input file:

```
$ go run go-csp-tp input/1.txt
```

Before running, please make sure that the package is located in GOPATH and the input path is specified correctly.

### Package

The main package consists of the following files:

* **main.go** runs the program
* **backtrack.go** implements the core function with recursive backtracking 
* **tiles.go** stores operations on tiles
* **helpers.go** includes helper functions to ease computation
* **io.go** takes care of input and output operations.
* **global.go** stores global declarations
