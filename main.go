package main

import "fmt"

// goroutine

func main() {

	go draw()
	go sum()
	go pathrecord()
	go sumrecord()
	ret := path(Point{0, 0})

	close(sumpoint)
	close(pathrecordpoint)
	wg.Wait()
	fmt.Println("ret:", ret)
}
