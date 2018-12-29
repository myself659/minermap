package main

import (
	"fmt"
	"os"
)

func pathrecordname() string {
	ret := fmt.Sprintf("record/path-%d-%d.txt", xylimit, sumlimit)
	return ret
}

func sumrecordname() string {
	ret := fmt.Sprintf("record/sum-%d-%d.txt", xylimit, sumlimit)
	return ret
}

func pathrecord() {
	fw, _ := os.Create(pathrecordname())
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
	fw, _ := os.Create(sumrecordname())
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
