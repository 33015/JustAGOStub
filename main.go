package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type csvreaderi interface {
	myReader(r io.Reader) [][]string
}
type DefaultReaderS struct{}

func (sa DefaultReaderS) myReader(r io.Reader) [][]string {
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return records
}

func ReadCSVFile(filename string, csvi csvreaderi) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	defer f.Close()

	records := csvi.myReader(f)
	return records
}

func main() {
	sa := &DefaultReaderS{}
	log.Println(ReadCSVFile("test.csv", sa))
}
