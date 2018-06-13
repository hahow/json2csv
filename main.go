package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
	"github.com/thoas/go-funk"
)

type LineReader interface {
	ReadBytes(delim byte) (line []byte, err error)
}

func main() {
	inputFile := flag.String("i", "", "/path/to/input.json (optional; default is stdin)")
	outputFile := flag.String("o", "", "/path/to/output.csv (optional; default is stdout)")
	outputDelim := flag.String("d", ",", "delimiter used for output values")
	showVersion := flag.Bool("version", false, "print version string")
	printHeader := flag.Bool("p", false, "prints header to output")
	keys := StringArray{}
	flag.Var(&keys, "k", "fields to output")
	flag.Parse()

	if *showVersion {
		fmt.Printf("json2csv %s\n", VERSION)
		return
	}

	var reader *bufio.Reader
	var writer *csv.Writer
	if *inputFile != "" {
		file, err := os.OpenFile(*inputFile, os.O_RDONLY, 0600)
		if err != nil {
			log.Printf("Error %s opening input file %v", err, *inputFile)
			os.Exit(1)
		}
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Printf("Error %s opening output file %v", err, *outputFile)
			os.Exit(1)
		}
		writer = csv.NewWriter(file)
	} else {
		writer = csv.NewWriter(os.Stdout)
	}

	delim, _ := utf8.DecodeRuneInString(*outputDelim)
	writer.Comma = delim

	json2csv(reader, writer, keys, *printHeader)
}

func get_value(data []interface{}, index int) string {
	if index < len(data) {
		return data[index].(string)
	}

	return ""
}

func json2csv(r LineReader, w *csv.Writer, keys []string, printHeader bool) {
	var line []byte
	var err error
	var index_of_keys []int
	line_count := 0

	for {
		if err == io.EOF {
			return
		}
		line, err = r.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Input ERROR: %s", err)
				break
			}
		}
		line_count++
		if (line_count == 1) {
			for _, key := range keys {
				var keys_data []interface{}
				err = json.Unmarshal(line, &keys_data)
				if err != nil {
					log.Printf("ERROR Decoding JSON at line %d: %s\n%s", line_count, err, line)
					continue
				}
				index := funk.IndexOf(keys_data, key)
				index_of_keys = append(index_of_keys, index)
			}
			continue
		}
		if len(line) == 0 {
			continue
		}

		if printHeader {
			w.Write(keys)
			w.Flush()
			printHeader = false
		}

		var data []interface{}
		err = json.Unmarshal(line, &data)
		if err != nil {
			log.Printf("ERROR Decoding JSON at line %d: %s\n%s", line_count, err, line)
			continue
		}

		var record []string
		for _, index := range index_of_keys {
			record = append(record, get_value(data, index))
		}

		w.Write(record)
		w.Flush()
	}
}
