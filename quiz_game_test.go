package main

import (
	"reflect"
	"strings"
	"testing"
)

type mockData string

func (m mockData) Read() *strings.Reader {
	return strings.NewReader(string(m))
}

func TestReadCSV(t *testing.T) {
	in := `5+5,10
7+3,10
1+1,2`
	data := mockData(in).Read()
	got := ReadCsv(data)
	want := [][]string{{"5+5", "10"}, {"7+3", "10"}, {"1+1", "2"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
