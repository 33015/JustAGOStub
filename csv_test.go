package main

import (
	"io"
	"reflect"
	"testing"
)

type TestReaderS struct{}

func (st TestReaderS) myReader(r io.Reader) [][]string {
	retvalue := [][]string{
		{"A", "B", "C"},
		{"4", "5", "6"},
	}
	return retvalue
}

func TestReadCSV(t *testing.T) {
	st := &TestReaderS{}
	retvalue := [][]string{
		{"A", "B", "C"},
		{"4", "5", "6"},
	}
	records := ReadCSVFile("test.csv", st)

	if reflect.DeepEqual(records, retvalue) == false {
		t.Error("records nicht gleich")
	}
}
