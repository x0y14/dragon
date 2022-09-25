package parser

import "dragon/common/tokenizer"

type Parser interface {
	Parse(*tokenizer.Token) ([]*Node, error)
}
