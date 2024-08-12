package xmas_bash_v1

type Option int

const (
	Rock Option = iota
	Paper
	Scissors
)

type HandResult int

const (
	Tie HandResult = iota
)

type Game struct {
}

func (game *Game) FirstPlayerSelect(option Option) {

}

func (game *Game) SecondPlayerSelect(option Option) {

}

func (game *Game) Result() HandResult {
	return Tie
}
