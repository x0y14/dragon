package tokenizer

type Token struct {
	Kind interface {
		int
		String() string
	}
	Pos *Position

	// 即値格納用フィールド
	S string
	I int64
	F float64

	Next *Token
}

func NewToken(kind interface {
	int
	String() string
}, pos *Position, s string, f float64, i int64) *Token {
	return &Token{
		Pos:  pos,
		Kind: kind,
		S:    s,
		I:    i,
		F:    f,
		Next: nil,
	}
}
