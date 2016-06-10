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

type yamlValue struct {
	underlying interface{}
}

func (value *yamlValue) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&value.underlying)
	if err != nil {
		return err
	}

	if _, found := value.underlying.(map[interface{}]interface{}); found {
		var bridge map[string]yamlValue
		converted := make(map[string]interface{}, len(bridge))
		err = unmarshal(&bridge)
		if err != nil {
			return err
		}

		for k, v := range bridge {
			converted[k] = v.underlying
		}

		value.underlying = converted
	}

	return nil
}

func (value yamlValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.underlying)
}

func main() {
	flag.Usage = func() {
		progname := os.Args[0]
		fmt.Printf("Usage of %s:\n", progname)
		fmt.Printf("\t%s [input [output]]\n", progname)
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

	var m yamlValue

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("Unable to parse input: %v\n", err)
	}

	data, err = json.Marshal(m)
	if err != nil {
		log.Fatalf("Unable to format output: %v\n", err)
	}

	_, err = output.Write(data)
	if err != nil {
		log.Fatalf("Unable to write output: %v\n", err)
	}
}
