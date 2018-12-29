package main

import "fmt"

func sum() {
	total := 0

	for {
		p, ok := <-sumpoint
		if ok {
			ps := p.SymmetryGroup()
			total = total + len(ps)
			for _, item := range ps {
				drawpoint <- item
				sumrecordpoint <- item
			}
		} else {
			fmt.Println("The total:", total)
			close(drawpoint)
			close(sumrecordpoint)
			return
		}
	}
}
