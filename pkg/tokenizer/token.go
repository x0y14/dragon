package tokenizer

type Token struct {
	Pos  *Position
	Kind interface {
		int
		String() string
	}
}
