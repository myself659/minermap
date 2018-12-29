package main

import (
	"fmt"
	"os"
)

func pathrecord() {
	fw, _ := os.Create("record/path.txt")
	wg.Add(1)
	defer wg.Done()
	i := 0
	for {
		p, ok := <-pathrecordpoint
		if ok {
			fmt.Fprintf(fw, "%d: %+v\n", i, p)
		} else {
			break
		}
		i++
	}
	fw.Close()
}

func sumrecord() {
	fw, _ := os.Create("record/sum.txt")
	wg.Add(1)
	defer wg.Done()
	i := 0
	for {
		p, ok := <-sumrecordpoint
		if ok {
			fmt.Fprintf(fw, "%d: %+v\n", i, p)
		} else {
			break
		}
		i++
	}
	fw.Close()
}
