package game

import (
	"reflect"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	want := [3][3]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	if got := createBoard(); !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, wanted: %v", got, want)
	}
}
