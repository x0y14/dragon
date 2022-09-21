package parser

import "dragon/pkg/tokenizer"

type Node struct {
	Pos  *tokenizer.Position
	Kind interface {
		int
		String() string
	}
}
