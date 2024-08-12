package xmas_bash_v1_test

import (
	"testing"
	xmas "xmas-bash/xmas_bash_v1"
)

func TestFullGame(t *testing.T) {
	game := xmas.Game{}

	game.FirstPlayerSelect(xmas.Rock)
	game.SecondPlayerSelect(xmas.Rock)
	expected := xmas.Tie
	if game.Result() != expected {
		t.Errorf("Result: Expected %d but got %d ", expected, game.Result())
	}
}
