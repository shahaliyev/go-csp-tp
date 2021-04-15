package main

// Global declarations
var landscape [][]int

var tiles = make(map[string]int)
var tileNames = [3]string{"OUTER_BOUNDARY", "EL_SHAPE", "FULL_BLOCK"}

var targets = make(map[int]int)


var graph = make(map[int][]int)
var domains = make(map[int][]int)

var used = make(map[int]int)

var colorCount int
var domainCount int

const INF = 0x3f3f3f3
