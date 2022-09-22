package tokenizer

import (
	"fmt"
)

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

// NewPositionFromString 入力の内容を読み取った後の時の位置を返す, Debug用
func NewPositionFromString(s string) *Position {
	p := &Position{
		LineNo:  1,
		LineAt:  0,
		WholeAt: 0,
	}
	for _, r := range []rune(s) {
		if r == '\n' {
			p.LineNo++
			p.WholeAt++
			p.LineAt = 0
			continue
		}
		p.WholeAt++
		p.LineAt++
	}
	return p
}

func (p *Position) Clone() *Position {
	return &Position{LineNo: p.LineNo, LineAt: p.LineAt, WholeAt: p.WholeAt}
}
