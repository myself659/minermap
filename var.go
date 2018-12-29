package main

import "sync"

var xylimit = 7
var sumlimit = 7
var pointVisit = make(map[Point]bool)
var drawpoint = make(chan Point)
var sumpoint = make(chan Point)
var wg sync.WaitGroup
var pathrecordpoint = make(chan Point)
var sumrecordpoint = make(chan Point)
