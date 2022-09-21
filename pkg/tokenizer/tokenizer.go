package tokenizer

import "dragon/pkg/scan"

type Tokenizer interface {
	Tokenize(*scan.ScriptFile) (*Token, error)
}
