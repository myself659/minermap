package main

import "flag"

func init() {
	flag.IntVar(&xylimit, "xylimit", 7, "limit for xy label")

	flag.IntVar(&sumlimit, "sumlimit", 7, "limit for sum")
	flag.Parse()
}
