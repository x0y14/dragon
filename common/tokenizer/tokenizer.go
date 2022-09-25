package tokenizer

type Tokenizer interface {
	Tokenize([]rune) (*Token, error)
}
