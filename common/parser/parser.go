package parser

import "github.com/x0y14/dragon/common/tokenizer"

type Parser interface {
	Parse(*tokenizer.Token) ([]*Node, error)
}
