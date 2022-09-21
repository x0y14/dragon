package tokenizer

import "dragon/pkg/scan"

type Tokenizer interface {
	Tokenize(scriptFile *scan.ScriptFile) (*Token, error)
}
