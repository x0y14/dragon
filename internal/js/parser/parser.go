package parser

import (
	"dragon/internal/js/tokenizer"
	commonParser "dragon/pkg/parser"
	commonTokenizer "dragon/pkg/tokenizer"
	"fmt"
)

type Parser struct {
	currentToken *commonTokenizer.Token
}

func (p *Parser) isEof() bool {
	return p.currentToken.Kind == tokenizer.Eof
}

func (p *Parser) consumeKind(k tokenizer.Kind) *commonTokenizer.Token {
	if p.currentToken.Kind == k {
		cur := p.currentToken.Clone()
		p.currentToken = p.currentToken.Next
		return cur
	}
	return nil
}

func (p *Parser) expectKind(k tokenizer.Kind) (*commonTokenizer.Token, error) {
	if p.currentToken.Kind != k {
		return nil, fmt.Errorf("unexpected token err: expected: %s, but found: %s", k.String(), p.currentToken.Kind.String())
	}
	cur := p.currentToken
	p.currentToken = p.currentToken.Next
	return cur, nil
}

func (p *Parser) parseFuncParams() (*commonParser.Node, error) {
	var params []*commonParser.Node
	for !p.isEof() {
		id, err := p.expectKind(tokenizer.Identifier)
		if err != nil {
			return nil, fmt.Errorf("failed to parse func params: %v", err)
		}
		params = append(params, NewIdentifierNode(id))
		// identの次が,だったらまだ続くということ
		// そうでないなら終わり
		if comma := p.consumeKind(tokenizer.SYComma); comma == nil {
			break
		}
	}
	return NewNodeWithOutImmediate(FuncParams, nil, nil, nil, params), nil
}

func (p *Parser) parseFuncDefine() (*commonParser.Node, error) {
	// function ident ( params ) stmt
	// function
	_, err := p.expectKind(tokenizer.KWFunction)
	if err != nil {
		return nil, fmt.Errorf("failed to parse func define: %v", err)
	}

	// ident
	id, err := p.expectKind(tokenizer.Identifier)
	if err != nil {
		return nil, fmt.Errorf("failed to parse func define: %v", err)
	}
	identifier := NewIdentifierNode(id)

	// (
	if _, err = p.expectKind(tokenizer.SYLrb); err != nil {
		return nil, fmt.Errorf("failed to parse func define: %v", err)
	}

	var params *commonParser.Node = nil
	// )でないのであれば、
	if rrb := p.consumeKind(tokenizer.SYRrb); rrb == nil {
		params, err = p.parseFuncParams()
		if err != nil {
			return nil, fmt.Errorf("failed to parse func define: %v", err)
		}
		if _, err = p.expectKind(tokenizer.SYRrb); err != nil {
			return nil, fmt.Errorf("failed to parse func define: %v", err)
		}
	}

	block, err := p.parseBlock()
	if err != nil {
		return nil, fmt.Errorf("failed to parse func define: %v", err)
	}

	// childrenに入れるべきか
	return NewNodeWithOutImmediate(FuncDefine, identifier, params, block, nil), nil
}

func (p *Parser) parseBlock() (*commonParser.Node, error) {
	_, err := p.expectKind(tokenizer.SYLcb)
	if err != nil {
		return nil, fmt.Errorf("failed to parse block: %v", err)
	}

	var statements []*commonParser.Node
	for p.consumeKind(tokenizer.SYRcb) == nil {
		stmt, err := p.statement()
		if err != nil {
			return nil, fmt.Errorf("failed to parse block: %v", err)
		}
		statements = append(statements, stmt)
	}
	// forの条件としてRCBは処分されている。
	return NewNodeWithOutImmediate(Block, nil, nil, nil, statements), nil
}

func (p *Parser) parseVarDeclareAndDefine() (*commonParser.Node, error) {
	// var ident;
	// var ident = v;

	// var
	if _, err := p.expectKind(tokenizer.KWVar); err != nil {
		return nil, fmt.Errorf("failed to parse varDeclare/varDefien: %v", err)
	}

	// ident
	id, err := p.expectKind(tokenizer.Identifier)
	if err != nil {
		return nil, fmt.Errorf("failed to parse varDeclare/varDefine: %v", err)
	}
	identifier := NewIdentifierNode(id)

	// =
	// assがnilならdeclare
	if ass := p.consumeKind(tokenizer.OPAssign); ass == nil {
		return NewNodeWithOutImmediate(VarDeclare, nil, nil, identifier, nil), nil
	}

	// value
	// assignは上で消費済み
	v, err := p.statement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse varDefine: %v", err)
	}

	return NewNodeWithOutImmediate(VarDefine, identifier, v, nil, nil), nil
}

func (p *Parser) parseLetDeclareAndDefine() (*commonParser.Node, error) {
	return nil, nil
}

func (p *Parser) parseConst() (*commonParser.Node, error) {
	// const ident = v;

	// const:
	if _, err := p.expectKind(tokenizer.KWConst); err != nil {
		return nil, fmt.Errorf("failed to parse const: %v", err)
	}

	// ident:
	id, err := p.expectKind(tokenizer.Identifier)
	if err != nil {
		return nil, fmt.Errorf("failed to parse const: %v", err)
	}
	identifier := NewIdentifierNode(id)

	// =
	if _, err = p.expectKind(tokenizer.OPAssign); err != nil {
		return nil, fmt.Errorf("faield to parse const: %v", err)
	}

	// value
	v, err := p.statement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse ")
	}

	return &commonParser.Node{
		Kind:     ConstDefine,
		LHS:      identifier,
		RHS:      v,
		Child:    nil,
		Children: nil,
	}, nil
}

func (p *Parser) statement() (*commonParser.Node, error) {
	var node *commonParser.Node
	var err error
	switch p.currentToken.Kind {
	default:
		node, err = p.expression()
		_ = p.consumeKind(tokenizer.SYSemi)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to parse statement: %v", err)
	}
	if node == nil {
		return nil, fmt.Errorf("failed to parse statetment: expected token: %v", p.currentToken.Kind.String())
	}
	return node, err
}

func (p *Parser) expression() (*commonParser.Node, error) {
	return p.assignment()
}

func (p *Parser) assignment() (*commonParser.Node, error) {
	switch p.currentToken.Kind {
	case tokenizer.KWVar:
		return p.parseVarDeclareAndDefine()
	case tokenizer.KWLet:
		return p.parseLetDeclareAndDefine()
	case tokenizer.KWConst:
		return p.parseConst()
	}
	return nil, nil
}

func (p *Parser) Parse(token *commonTokenizer.Token) ([]*commonParser.Node, error) {
	p.currentToken = token
	var nodes []*commonParser.Node

	for !p.isEof() {
		nd, err := p.statement()
		if err != nil {
			return nil, fmt.Errorf("failed to parse token: %v", err)
		}
		nodes = append(nodes, nd)
	}

	return nodes, nil
}
