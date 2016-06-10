package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	flag.Usage = func() {
		progname := os.Args[0]
		fmt.Printf("Usage of %s:\n", progname)
		fmt.Printf("\t%s [input [output]]", progname)
		flag.PrintDefaults()
	}

	flag.Parse()

	input := os.Stdin
	output := os.Stdout
	var err error

	if flag.NArg() > 0 {
		inputFile := flag.Arg(0)
		input, err = os.Open(inputFile)
		if err != nil {
			log.Fatalf("Unable to open file %v: %v\n", inputFile, err)
		}
	}

	if flag.NArg() > 1 {
		outputFile := flag.Arg(1)
		output, err = os.OpenFile(outputFile, os.O_WRONLY, 0)
		if err != nil {
			log.Fatalf("Unable to open file %v: %v\n", outputFile, err)
		}
	}

	data, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatalf("Unable to read input: %v\n", err)
	}

	var m interface{}

	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("Unable to parse input: %v\n", err)
	}

	data, err = yaml.Marshal(m)
	if err != nil {
		log.Fatalf("Unable to format output: %v\n", err)
	}

	_, err = output.Write(data)
	if err != nil {
		log.Fatalf("Unable to write output: %v\n", err)
	}
}
