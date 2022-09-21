package tokenizer

type Position struct {
	LineNo  int
	LineAt  int
	WholeAt int
}

func NewPosition(lno, lat, wat int) *Position {
	return &Position{
		LineNo:  lno,
		LineAt:  lat,
		WholeAt: wat,
	}
}
