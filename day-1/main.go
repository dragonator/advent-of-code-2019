package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dragonator/advent-of-code-2019/day-1/pkg"
)

func main() {
	inputFile := flag.String("input-file", "", "Puzzle input file")
	part := flag.String("part", "", "Part of the day to run (1 or 2)")
	flag.Parse()

	if inputFile == nil || *inputFile == "" {
		fmt.Println("error: '-input-file' flag is missing or value is empty string")
		flag.Usage()
		os.Exit(1)
	}

	if part == nil || (*part != "1" && *part != "2") {
		fmt.Println("error: '-part' flag is missing or has unexpected value")
		flag.Usage()
		os.Exit(1)
	}

	_, err := os.Stat(*inputFile)
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("error: provided path for input-file is invalid or does not exist: %s\n", err.Error())
		flag.Usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("error: an error occured when tried to read the file: %s\n", err.Error())
		os.Exit(1)
	}

	var modulesOnly bool
	if *part == "1" {
		modulesOnly = true
	} else { // part == "2"
		modulesOnly = false
	}
	moduleMasses := strings.Split(string(data), "\n")
	res, err := pkg.CalculateTotalFuel(moduleMasses, modulesOnly)
	if err != nil {
		fmt.Printf("error: failed to calculate total fuel required: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Total fuel: %d\n", res)
}
