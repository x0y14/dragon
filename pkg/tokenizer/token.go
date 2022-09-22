package tokenizer

// TokenKind String() stringを持ったint型を想定してる
// String() -> "Eof"の場合があるようにしてください
type TokenKind interface {
	String() string
}

type Token struct {
	Kind TokenKind
	Pos  *Position

	// 即値格納用フィールド
	S string
	I int64
	F float64

	Next *Token
}

func NewToken(kind TokenKind, pos *Position, s string, f float64, i int64) *Token {
	return &Token{
		Pos:  pos,
		Kind: kind,
		S:    s,
		I:    i,
		F:    f,
		Next: nil,
	}
}

func (t *Token) IsEof() bool {
	return t.Kind.String() == "Eof"
}
