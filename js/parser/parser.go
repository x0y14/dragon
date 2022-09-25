package parser

import (
	commonParser "dragon/common/parser"
	commonTokenizer "dragon/common/tokenizer"
	"dragon/js/tokenizer"
	"fmt"
)

type Parser struct {
	currentToken *commonTokenizer.Token
}

func NewParser() *Parser {
	return &Parser{}
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

func (p *Parser) consume() *commonTokenizer.Token {
	cur := p.currentToken.Clone()
	p.currentToken = p.currentToken.Next
	return cur
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
	return NewNodeWithOutBothSidesAndImmediate(FuncParams, params, nil, nil, nil), nil
}

func (p *Parser) parseCallArgs() (*commonParser.Node, error) {
	var args []*commonParser.Node
	for !p.isEof() && p.currentToken.Kind != tokenizer.SYRrb {
		arg, err := p.expression()
		if err != nil {
			return nil, fmt.Errorf("failed to parse call-args: %v", err)
		}
		args = append(args, arg)
		if p.currentToken.Kind == tokenizer.SYComma {
			_ = p.consume()
		} else {
			break
		}
	}
	return NewNodeWithOutBothSidesAndImmediate(CallArgs, args, nil, nil, nil), nil
}

// stmt階層
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
	return NewNodeWithOutBothSidesAndImmediate(FuncDefine, nil, identifier, params, block), nil
}

func (p *Parser) parseReturn() (*commonParser.Node, error) {
	_, err := p.expectKind(tokenizer.KWReturn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse return: %v", err)
	}

	// ";", "}"であるならば
	switch p.currentToken.Kind {
	case tokenizer.SYSemi, tokenizer.SYRcb:
		// 何も要素を返却しない
		return NewNodeWithOutBothSidesAndImmediate(Return, nil, nil, nil, nil), nil
	default:
		v, err := p.expression()
		if err != nil {
			return nil, fmt.Errorf("failed to parse return: %v", err)
		}
		return NewNodeWithOutBothSidesAndImmediate(Return, nil, v, nil, nil), nil
	}
}

func (p *Parser) parseIfElse() (*commonParser.Node, error) {
	// if "(" expr? ")" STMT1 ("else" STMT2)?
	// if
	if _, err := p.expectKind(tokenizer.KWIf); err != nil {
		return nil, fmt.Errorf("failed to parse if: %v", err)
	}
	// "("
	if _, err := p.expectKind(tokenizer.SYLrb); err != nil {
		return nil, fmt.Errorf("failed to parse if: %v", err)
	}
	// expr
	cond, err := p.expression()
	if err != nil {
		return nil, fmt.Errorf("failed to parse if: %v", err)
	}
	// ")"
	if _, err := p.expectKind(tokenizer.SYRrb); err != nil {
		return nil, fmt.Errorf("failed to parse if: %v", err)
	}

	// stmt1
	ifBlock, err := p.statement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse if: %v", err)
	}

	if els := p.consumeKind(tokenizer.KWElse); els != nil {
		// Kind: if-else
		// stmt2
		elseBlock, err := p.statement()
		if err != nil {
			return nil, fmt.Errorf("failed to parse if-else: %v", err)
		}
		return NewNodeWithOutBothSidesAndImmediate(IfElse, nil, cond, ifBlock, elseBlock), nil
	}
	// Kind: if
	return NewNodeWithOutBothSidesAndImmediate(If, nil, cond, ifBlock, nil), nil
}

func (p *Parser) parseWhile() (*commonParser.Node, error) {
	return nil, nil
}

func (p *Parser) parseFor() (*commonParser.Node, error) {
	return nil, nil
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
	// for文の条件としてRCBは処分されている。
	return NewNodeWithOutBothSidesAndImmediate(Block, statements, nil, nil, nil), nil
}

// assign階層
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
		return NewNodeWithOutBothSidesAndImmediate(VarDeclare, nil, identifier, nil, nil), nil
	}

	// value
	// assignは上で消費済み
	v, err := p.statement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse varDefine: %v", err)
	}

	return NewNodeWithOutImmediate(VarDefine, identifier, v, nil, nil, nil, nil), nil
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

	return NewNodeWithOutImmediate(ConstDefine, identifier, v, nil, nil, nil, nil), nil
}

// 各階層
func (p *Parser) statement() (*commonParser.Node, error) {
	var node *commonParser.Node
	var err error
	switch p.currentToken.Kind {
	case tokenizer.KWFunction:
		node, err = p.parseFuncDefine()
	case tokenizer.KWReturn:
		node, err = p.parseReturn()
		_ = p.consumeKind(tokenizer.SYSemi)
	case tokenizer.KWIf:
		node, err = p.parseIfElse()
	case tokenizer.KWWhile:
		node, err = p.parseWhile()
	case tokenizer.KWFor:
		node, err = p.parseFor()
	case tokenizer.SYLcb:
		node, err = p.parseBlock()
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
	default:
		return p.andor()
	}
}

func (p *Parser) andor() (*commonParser.Node, error) {
	nd, err := p.equality()
	if err != nil {
		return nil, fmt.Errorf("failed to parse andor: %v", err)
	}
	for {
		switch p.currentToken.Kind {
		case And:
			rhsNd, err := p.equality()
			if err != nil {
				return nil, fmt.Errorf("failed to parse andor: %v", err)
			}
			nd = NewNodeBothSides(And, nd, rhsNd)
		case Or:
			rhsNd, err := p.equality()
			if err != nil {
				return nil, fmt.Errorf("failed to parse andor: %v", err)
			}
			nd = NewNodeBothSides(Or, nd, rhsNd)
		default:
			return nd, nil
		}
	}
}

func (p *Parser) equality() (*commonParser.Node, error) {
	nd, err := p.relational()
	if err != nil {
		return nil, fmt.Errorf("failed to parse equality: %v", err)
	}
	for {
		switch p.currentToken.Kind {
		case tokenizer.LEq:
			rhsNd, err := p.relational()
			if err != nil {
				return nil, fmt.Errorf("failed to parse equality: %v", err)
			}
			nd = NewNodeBothSides(Eq, nd, rhsNd)
		case tokenizer.LNe:
			rhsNd, err := p.relational()
			if err != nil {
				return nil, fmt.Errorf("failed to parse equality: %v", err)
			}
			nd = NewNodeBothSides(Ne, nd, rhsNd)
		default:
			return nd, nil
		}
	}
}

func (p *Parser) relational() (*commonParser.Node, error) {
	nd, err := p.add()
	if err != nil {
		return nil, fmt.Errorf("failed to parse relational: %v", err)
	}
	for {
		switch p.currentToken.Kind {
		case tokenizer.LLt:
			rhsNd, err := p.add()
			if err != nil {
				return nil, fmt.Errorf("faield to parse relational: %v", err)
			}
			nd = NewNodeBothSides(Lt, nd, rhsNd)
		case tokenizer.LLe:
			rhsNd, err := p.add()
			if err != nil {
				return nil, fmt.Errorf("faield to parse relational: %v", err)
			}
			nd = NewNodeBothSides(Le, nd, rhsNd)
		case tokenizer.LGt:
			rhsNd, err := p.add()
			if err != nil {
				return nil, fmt.Errorf("faield to parse relational: %v", err)
			}
			nd = NewNodeBothSides(Gt, nd, rhsNd)
		case tokenizer.LGe:
			rhsNd, err := p.add()
			if err != nil {
				return nil, fmt.Errorf("faield to parse relational: %v", err)
			}
			nd = NewNodeBothSides(Ge, nd, rhsNd)
		default:
			return nd, nil
		}
	}
}

func (p *Parser) add() (*commonParser.Node, error) {
	nd, err := p.mul()
	if err != nil {
		return nil, fmt.Errorf("failed to parse add: %v", err)
	}
	for {
		switch p.currentToken.Kind {
		case tokenizer.OPAdd:
			rhsNd, err := p.mul()
			if err != nil {
				return nil, fmt.Errorf("failed to parse add: %v", err)
			}
			nd = NewNodeBothSides(Add, nd, rhsNd)
		case tokenizer.OPSub:
			rhsNd, err := p.mul()
			if err != nil {
				return nil, fmt.Errorf("failed to parse add: %v", err)
			}
			nd = NewNodeBothSides(Sub, nd, rhsNd)
		default:
			return nd, nil
		}
	}
}

func (p *Parser) mul() (*commonParser.Node, error) {
	nd, err := p.unary()
	if err != nil {
		return nil, fmt.Errorf("failed to parse mul: %v", err)
	}
	for {
		switch p.currentToken.Kind {
		case Mul:
			rhsNd, err := p.unary()
			if err != nil {
				return nil, fmt.Errorf("failed to parse mul: %v", err)
			}
			nd = NewNodeBothSides(Mul, nd, rhsNd)
		case Div:
			rhsNd, err := p.unary()
			if err != nil {
				return nil, fmt.Errorf("failed to parse mul: %v", err)
			}
			nd = NewNodeBothSides(Div, nd, rhsNd)
		case Mod:
			rhsNd, err := p.unary()
			if err != nil {
				return nil, fmt.Errorf("failed to parse mul: %v", err)
			}
			nd = NewNodeBothSides(Mod, nd, rhsNd)
		default:
			return nd, nil
		}
	}
}

func (p *Parser) unary() (*commonParser.Node, error) {
	switch p.currentToken.Kind {
	case tokenizer.OPAdd:
		_ = p.consumeKind(tokenizer.OPAdd)
		return p.primary()
	case tokenizer.OPSub:
		_ = p.consumeKind(tokenizer.OPSub)
		rhsNd, err := p.primary()
		if err != nil {
			return nil, fmt.Errorf("faield to parse unary: %v", err)
		}
		lhsNd := commonParser.NewNode(NumberInteger, nil, nil, nil, nil, nil, nil, "", 0, 0, false)
		return NewNodeBothSides(Sub, lhsNd, rhsNd), nil
	case tokenizer.LNot:
		_ = p.consumeKind(tokenizer.LNot)
		v, err := p.primary()
		if err != nil {
			return nil, fmt.Errorf("faile to parse unary: %v", err)
		}
		return NewNodeWithOutBothSidesAndImmediate(Not, nil, v, nil, nil), nil
	default:
		return p.primary()
	}
}

func (p *Parser) primary() (*commonParser.Node, error) {
	return p.access()
}

func (p *Parser) access() (*commonParser.Node, error) {
	nd, err := p.literal()
	if err != nil {
		return nil, fmt.Errorf("failed to parse access: %v", err)
	}
	for {
		switch p.currentToken.Kind {
		case tokenizer.SYLsb:
			// [
			_ = p.consumeKind(tokenizer.SYLsb)
			index, err := p.expression()
			if err != nil {
				return nil, fmt.Errorf("failed to parse access: %v", err)
			}
			// ]
			if _, err = p.expectKind(tokenizer.SYRsb); err != nil {
				return nil, fmt.Errorf("failed to parse access: %v", err)
			}
			nd = NewNodeBothSides(ObjectAccess, nd, index)
		case tokenizer.SYDot:
			_ = p.consumeKind(tokenizer.SYDot)
			rhsNd, err := p.primary()
			if err != nil {
				return nil, fmt.Errorf("failed to parse access: %v", err)
			}
			nd = NewNodeBothSides(Field, nd, rhsNd)
		default:
			return nd, nil
		}
	}
}

func (p *Parser) literal() (*commonParser.Node, error) {
	switch p.currentToken.Kind {
	case tokenizer.SYLrb:
		_ = p.consumeKind(tokenizer.SYLrb)
		expr, err := p.expression()
		if err != nil {
			return nil, fmt.Errorf("failed to parse literal: %v", err)
		}
		_, err = p.expectKind(tokenizer.SYRrb)
		if err != nil {
			return nil, fmt.Errorf("failed to parse literal: %v", err)
		}
		return NewNodeWithOutBothSidesAndImmediate(Parenthesis, nil, expr, nil, nil), nil
	case tokenizer.Identifier:
		id := p.consumeKind(tokenizer.Identifier)
		ident := NewIdentifierNode(id)

		if lrb := p.consumeKind(tokenizer.SYLrb); lrb != nil {
			// call
			args, err := p.parseCallArgs()
			if err != nil {
				return nil, fmt.Errorf("failed to parse literal: %v", err)
			}
			if _, err := p.expectKind(tokenizer.SYRrb); err != nil {
				return nil, fmt.Errorf("faield to parse literal: %v", err)
			}
			return NewNodeWithOutBothSidesAndImmediate(Call, nil, ident, args, nil), nil
		}
		return ident, nil
	case tokenizer.String, tokenizer.Decimal, tokenizer.Integer, tokenizer.KWTrue, tokenizer.KWFalse, tokenizer.KWNull:
		return NewImmediateNode(p.consume())
	case tokenizer.SYLsb: // array
	case tokenizer.SYLcb: // object
	}
	return nil, fmt.Errorf("unimpleremted: %s", p.currentToken.Kind.String())
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
