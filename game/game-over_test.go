package game

import "testing"

func TestIsWinner(t *testing.T) {
	tests := []struct {
		board [3][3]string
		want  bool
	}{
		{[3][3]string{
			{"o", "o", "o"},
			{"x", "x", "."},
			{".", ".", "."},
		}, true},
		{[3][3]string{
			{"o", "x", "."},
			{"o", "x", "."},
			{"o", ".", "."},
		}, true},
		{[3][3]string{
			{"o", ".", "."},
			{"x", "o", "."},
			{".", "x", "o"},
		}, true},
		{[3][3]string{
			{"o", ".", "o"},
			{".", "x", "."},
			{"x", ".", "o"},
		}, false},
	}

	for _, test := range tests {
		if got := isWinner(test.board); got != test.want {
			t.Errorf("got: %v, wanted: %v", got, test.want)
		}
	}
}
