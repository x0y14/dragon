package tokenizer

import "dragon/pkg/scan"

type Tokenizer interface {
	Tokenize(script *scan.ScriptFile) *Token
}
