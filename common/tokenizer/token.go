package tokenizer

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

func (t *Token) IsEof() bool {
	return t.Kind.String() == "Eof"
}

func (t *Token) Clone() *Token {
	token_ := NewToken(t.Kind, t.Pos, t.S, t.F, t.I)
	token_.Next = t.Next
	return token_
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
