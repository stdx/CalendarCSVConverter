package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Event struct {
}

// https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if csvWasPiped(fi) {
		err = tryProcessInputCsv(bufio.NewReader(os.Stdin))
	} else {
		flag.Parse()
		if flag.NArg() != 1 {
			err = errors.New("Need exactly one input")
		} else {
			err = tryProcessInputCsvFromFileArg(flag.Arg(0))
		}
	}

	if err != nil {
		panic(err)
	}

}

func csvWasPiped(fi os.FileInfo) bool {
	return fi.Mode()&os.ModeNamedPipe == 0
}

func tryProcessInputCsvFromFileArg(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed reading file")
		return nil
	}
	defer f.Close()

	return tryProcessInputCsv(bufio.NewReader(f))
}

func tryProcessInputCsv(r *bufio.Reader) error {
	csvr := csv.NewReader(r)
	csvr.Comma = ';'
	csvr.Comment = '#'

	events := []*Event{}
	for {
		csvLine, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		e := toEvent(csvLine)
		events = append(events, e)
	}

	// TODO do something with the events


	return nil
}

func toEvent(csvLine []string) *Event {
	return &Event{};
}
