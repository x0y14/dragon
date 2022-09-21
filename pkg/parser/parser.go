package parser

import "dragon/pkg/tokenizer"

type Parser interface {
	Parse(*tokenizer.Token) ([]*Node, error)
}
