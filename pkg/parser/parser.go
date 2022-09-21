package parser

import "dragon/pkg/tokenizer"

type Parser interface {
	Parse(tokenChain *tokenizer.Token) ([]*Node, error)
}
