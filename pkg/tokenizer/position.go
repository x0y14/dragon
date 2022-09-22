package tokenizer

import "fmt"

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

func (p *Position) String() string {
	return fmt.Sprintf("%s:%s", p.LineNo, p.LineAt)
}
