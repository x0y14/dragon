package parser

import (
	commonParser "dragon/common/parser"
	commonTokenizer "dragon/common/tokenizer"
	"dragon/html/tokenizer"
	"fmt"
	"log"
	"strings"
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

func (p *Parser) isTagName(tok *commonTokenizer.Token, tagName string) bool {
	if tok.Kind != tokenizer.Text {
		return false
	}

	return strings.ToLower(tok.S) == strings.ToLower(tagName)
}

func (p *Parser) text() (string, error) {
	comment := ""
	for !p.isEof() {
		if ws := p.consumeKind(tokenizer.Whitespace); ws != nil {
			comment += " "
			continue
		}
		if txt := p.consumeKind(tokenizer.Text); txt != nil {
			comment += txt.S
			continue
		}
		break
	}
	return comment, nil
}

func (p *Parser) decl() (*commonParser.Node, error) {
	// comment or doctype

	// comment
	if hyp := p.consumeKind(tokenizer.Hyphen); hyp != nil {
		if _, err := p.expectKind(tokenizer.Hyphen); err != nil {
			return nil, err
		}
		comment, err := p.text()
		if err != nil {
			return nil, fmt.Errorf("failed to parse comment: %v", err)
		}

		if _, err = p.expectKind(tokenizer.Hyphen); err != nil {
			return nil, err
		}
		if _, err = p.expectKind(tokenizer.Hyphen); err != nil {
			return nil, err
		}
		if _, err = p.expectKind(tokenizer.TagEnd); err != nil {
			return nil, err
		}

		return commonParser.NewNode(Comment,
			nil, nil, nil,
			NewImmediateNode(String, comment, 0, 0, false), nil, nil, "", 0, 0, false), nil
	}

	if doctype := p.consumeKind(tokenizer.Text); doctype != nil && p.isTagName(doctype, "doctype") {
		if _, err := p.expectKind(tokenizer.Whitespace); err != nil {
			return nil, fmt.Errorf("failed to parse decl: %v", err)
		}
		documentType := p.consumeKind(tokenizer.Text)

		if _, err := p.expectKind(tokenizer.TagEnd); err != nil {
			return nil, err
		}

		return commonParser.NewNode(Doctype,
			nil, nil, nil,
			NewImmediateNode(String, documentType.S, 0, 0, false), nil, nil,
			"", 0, 0, false), nil
	}

	return nil, fmt.Errorf("unexpected token: %v", p.currentToken)
}

func (p *Parser) parameter() (*commonParser.Node, error) {
	ident, err := p.expectKind(tokenizer.Text)
	if err != nil {
		return nil, fmt.Errorf("faield to parse parameter: %v", err)
	}
	_ = p.consumeKind(tokenizer.Whitespace)
	if _, err = p.expectKind(tokenizer.Assign); err != nil {
		return nil, fmt.Errorf("faield to parse parameter: %v", err)
	}
	_ = p.consumeKind(tokenizer.Whitespace)
	value, err := p.expectKind(tokenizer.String)
	if err != nil {
		return nil, fmt.Errorf("faield to parse parameter: %v", err)
	}
	return commonParser.NewNode(Parameter,
		NewImmediateNode(Identifier, ident.S, 0, 0, false), NewImmediateNode(String, value.S, 0, 0, false), nil,
		nil, nil, nil,
		"", 0, 0, false), nil
}

func (p *Parser) parameters() (*commonParser.Node, error) {
	var params []*commonParser.Node
	// >がきたらタグ終了
	// />がきたらタグ終了
	for p.currentToken.Kind != tokenizer.TagEnd && p.currentToken.Kind != tokenizer.Slash {
		_ = p.consumeKind(tokenizer.Whitespace)
		param, err := p.parameter()
		if err != nil {
			return nil, fmt.Errorf("failed to parse parameters: %v", err)
		}
		params = append(params, param)
		_ = p.consumeKind(tokenizer.Whitespace)
	}

	if len(params) == 0 {
		return nil, nil
	}
	return commonParser.NewNode(Parameters,
		nil, nil, params,
		nil, nil, nil,
		"", 0, 0, false), nil
}

func (p *Parser) tag() (*commonParser.Node, error) {
	// already consume "<" by parse Tag

	//if sla := p.consumeKind(tokenizer.Slash); sla != nil {
	//	return nil, nil
	//}
	if p.currentToken.Kind == tokenizer.Slash {
		return nil, nil
	}

	// オープンタグ
	tagName, err := p.expectKind(tokenizer.Text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse open tag: %v", err)
	}

	_ = p.consumeKind(tokenizer.Whitespace)

	params, err := p.parameters()
	if err != nil {
		return nil, fmt.Errorf("failed to parse tag params: %v", err)
	}

	// /があればimgなどの自己完結タグ
	if sla := p.consumeKind(tokenizer.Slash); sla != nil {
		if _, err := p.expectKind(tokenizer.TagEnd); err != nil {
			return nil, fmt.Errorf("failed to parse solo-tag: %v", err)
		}
		return commonParser.NewNode(SoloTag,
			nil, nil, nil,
			NewImmediateNode(Identifier, tagName.S, 0, 0, false), params, nil,
			"", 0, 0, false), nil
	}

	// >
	if _, err := p.expectKind(tokenizer.TagEnd); err != nil {
		return nil, err
	}
	// オープンタグ終了

	var children []*commonParser.Node
	c, err := p.parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse children: %v", err)
	}
	children = c

	// クローズタグ
	// <
	//if _, err := p.expectKind(tokenizer.TagBegin); err != nil {
	//	return nil, fmt.Errorf("failed to parse close tag: %v", err)
	//}
	// /
	if _, err := p.expectKind(tokenizer.Slash); err != nil {
		return nil, fmt.Errorf("faield to parse close tag: %v", err)
	}

	closeTagName, err := p.expectKind(tokenizer.Text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse close tag: %v", err)
	}
	if _, err := p.expectKind(tokenizer.TagEnd); err != nil {
		return nil, fmt.Errorf("failed to parse close tag: %v", err)
	}

	// オープンとクローズでタグ名が変化していることはありえない
	if strings.ToLower(tagName.S) != strings.ToLower(closeTagName.S) {
		return nil, fmt.Errorf("open & close tag name is not match (open: %v, close: %v)", tagName.S, closeTagName.S)
	}

	return commonParser.NewNode(Tag,
		nil, nil, children,
		NewImmediateNode(Identifier, tagName.S, 0, 0, false), params, nil,
		"", 0, 0, false), nil
}

func (p *Parser) parseTag() (*commonParser.Node, error) {
	if excl := p.consumeKind(tokenizer.Exclamation); excl != nil {
		return p.decl()
	}
	return p.tag()
}

func (p *Parser) parse() ([]*commonParser.Node, error) {
	var nodes []*commonParser.Node
	for !p.isEof() {
		log.Printf("%v: %v", p.currentToken.Pos.String(), p.currentToken.Kind.String())
		if p.currentToken.Kind == tokenizer.Slash {
			break
		}

		if lt := p.consumeKind(tokenizer.TagBegin); lt != nil {
			nd, err := p.parseTag()
			if err != nil {
				return nil, fmt.Errorf("failed to parse tag: %v", err)
			}
			if nd != nil {
				nodes = append(nodes, nd)
			}
			continue
		}

		txt, err := p.text()
		if err != nil {
			return nil, fmt.Errorf("failed to parse text: %v", err)
		}
		nd := commonParser.NewNode(Text,
			nil, nil, nil,
			nil, nil, nil,
			txt, 0, 0, false)
		nodes = append(nodes, nd)
	}
	return nodes, nil
}

func (p *Parser) Parse(tok *commonTokenizer.Token) ([]*commonParser.Node, error) {
	p.currentToken = tok
	return p.parse()
}
