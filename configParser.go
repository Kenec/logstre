package main

import "flag"

var config = flag.String("config", "", "configuration file")
var help = flag.String("help", "", "")

func parseFlag() {
	flag.Parse()
}


