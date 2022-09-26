package tokenizer

import (
	"dragon/common/tokenizer"
	"fmt"
	"log"
	"strconv"
)

type Tokenizer struct {
	target     []rune
	currentPos *tokenizer.Position
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) isEof() bool {
	return t.currentPos.WholeAt >= len(t.target)
}

func (t *Tokenizer) peek(n int) rune {
	if t.currentPos.WholeAt+n >= len(t.target) {
		log.Fatalf("target item count is %d, so we can't access index %d", len(t.target), t.currentPos.WholeAt+n)
	}
	return t.target[t.currentPos.WholeAt+n]
}

func (t *Tokenizer) startWith(s string) bool {
	sRunes := []rune(s)
	for i := 0; i < len(sRunes); i++ {
		if t.peek(i) != sRunes[i] {
			return false
		}
	}
	return true
}

func (t *Tokenizer) currentRune() rune {
	return t.target[t.currentPos.WholeAt]
}

func (t *Tokenizer) moveHorizon(n int) *tokenizer.Position {
	t.currentPos.WholeAt += n
	t.currentPos.LineAt += n
	return t.currentPos
}

func (t *Tokenizer) moveNewline(n int) *tokenizer.Position {
	t.currentPos.WholeAt += n
	t.currentPos.LineNo += n
	t.currentPos.LineAt = 0
	return t.currentPos
}

func (t *Tokenizer) isWhite(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func (t *Tokenizer) isSingleSymbol() (string, bool) {
	for _, symbol := range []string{
		"<", ">", "!", "=", "-", "/", "&",
	} {
		if t.startWith(symbol) {
			return symbol, true
		}
	}
	return "", false
}

func (t *Tokenizer) isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

func (t *Tokenizer) isAlphabet(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func (t *Tokenizer) consumeWhite() string {
	var w string
	for !t.isEof() {
		if cr := t.currentRune(); t.isWhite(cr) && cr != '\n' {
			w += string(cr)
			t.moveHorizon(1)
		} else if cr == '\n' {
			w += string(cr)
			t.moveNewline(1)
		} else {
			break
		}
	}
	return w
}

func (t *Tokenizer) consumeIdent() string {
	var s string

	for !t.isEof() {
		if t.isAlphabet(t.currentRune()) || t.isNumber(t.currentRune()) || t.currentRune() == '_' {
			s += string(t.currentRune())
			t.moveHorizon(1)
		} else {
			break
		}
	}

	return s
}

func (t *Tokenizer) consumeNumber() (string, bool) {
	isDotIncluded := false
	var n string
	for !t.isEof() {
		if cr := t.currentRune(); t.isNumber(cr) {
			n += string(cr)
			t.moveHorizon(1)
			continue
		} else if t.currentRune() == '.' && t.isNumber(t.peek(1)) {
			// 今が.で、次が数字であれば...
			// そうでないならメンバだと思う...?
			n += string(t.currentRune())
			t.moveHorizon(1)
			isDotIncluded = true
			continue
		} else {
			break
		}
	}
	return n, isDotIncluded
}

func (t *Tokenizer) consumeString() string {
	var s string
	isSingle := false

	if t.currentRune() == '\'' {
		isSingle = true
	}

	t.moveHorizon(1) // consume start single/double quotation

	for !t.isEof() {
		if isSingle && t.currentRune() == '\'' {
			break
		} else if !isSingle && t.currentRune() == '"' {
			break
		}

		// escape
		if t.currentRune() == '\\' {
			if t.peek(1) == '"' {
				s += "\""
				t.moveHorizon(2)
				continue
			}
			if t.peek(1) == '\'' {
				s += "'"
				t.moveHorizon(2)
				continue
			}
			if t.peek(1) == '\\' {
				s += "\\"
				t.moveHorizon(2)
				continue
			}
			if t.peek(1) == 't' {
				s += "\t"
				t.moveHorizon(2)
				continue
			}
			if t.peek(1) == 'n' {
				s += "\n"
				t.moveHorizon(2)
				continue
			}
		}

		s += string(t.currentRune())
		t.moveHorizon(1)
	}

	t.moveHorizon(1) // consume end single/double quotation
	return s
}

func (t *Tokenizer) linkNewWhitespaceToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position, s string) *tokenizer.Token {
	tok := tokenizer.NewToken(Whitespace, startedAt, s, 0, 0)
	currentToken.Next = tok
	return tok
}

func (t *Tokenizer) linkNewSymbolToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position, symbol string) (*tokenizer.Token, error) {
	kind := symbolKind(symbol)
	if kind == Illegal {
		return nil, fmt.Errorf("unsupported symbol: %s", symbol)
	}
	tok := tokenizer.NewToken(kind, startedAt, "", 0, 0)
	currentToken.Next = tok
	return tok, nil
}

func (t *Tokenizer) linkNewDecimalToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position, n float64) *tokenizer.Token {
	tok := tokenizer.NewToken(Decimal, startedAt, "", n, 0)
	currentToken.Next = tok
	return tok
}

func (t *Tokenizer) linkNewStringToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position, s string) *tokenizer.Token {
	tok := tokenizer.NewToken(String, startedAt, s, 0, 0)
	currentToken.Next = tok
	return tok
}

func (t *Tokenizer) linkNewIdentToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position, ident string) *tokenizer.Token {
	kind := identKind(ident)
	tok := tokenizer.NewToken(kind, startedAt, ident, 0, 0)
	currentToken.Next = tok
	return tok
}

func (t *Tokenizer) linkNewIntegerToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position, n int64) *tokenizer.Token {
	tok := tokenizer.NewToken(Integer, startedAt, "", 0, n)
	currentToken.Next = tok
	return tok
}

func (t *Tokenizer) linkNewEofToken(currentToken *tokenizer.Token, startedAt *tokenizer.Position) *tokenizer.Token {
	tok := tokenizer.NewToken(Eof, startedAt, "", 0, 0)
	currentToken.Next = tok
	return tok
}

func (t *Tokenizer) Tokenize(target []rune) (*tokenizer.Token, error) {
	t.currentPos = tokenizer.NewPosition(1, 0, 0)
	t.target = target

	var head tokenizer.Token
	cur := &head

	for !t.isEof() {
		if symbol, ok := t.isSingleSymbol(); ok {
			var err error
			cur, err = t.linkNewSymbolToken(cur, t.currentPos.Clone(), symbol)
			if err != nil {
				return nil, err
			}
			// isSingleSymbol->startWithで場所を動かしてないので、動かしてあげる
			t.moveHorizon(len(symbol))
			continue
		}
		if t.isWhite(t.currentRune()) {
			startedAt := t.currentPos.Clone()
			ws := t.consumeWhite()
			cur = t.linkNewWhitespaceToken(cur, startedAt, ws)
			continue
		}
		// 数字
		if t.isNumber(t.currentRune()) {
			// consume関数で位置を移動するので、開始地点を保存しておく
			startedAt := t.currentPos.Clone()
			numStr, dotIncluded := t.consumeNumber()
			if dotIncluded {
				// 少数
				n, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return nil, fmt.Errorf("failed to parse float: %v", err)
				}
				cur = t.linkNewDecimalToken(cur, startedAt, n)
				// consumeで移動済み
				continue
			} else {
				// 整数
				n, err := strconv.ParseInt(numStr, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("failed to parse int: %v", err)
				}
				cur = t.linkNewIntegerToken(cur, startedAt, n)
				// consumeで移動済み
				continue
			}
		}
		// 文字列
		if t.currentRune() == '"' || t.currentRune() == '\'' {
			startedAt := t.currentPos.Clone()
			str := t.consumeString()
			cur = t.linkNewStringToken(cur, startedAt, str)
			// consumeで移動済み
			continue
		}
		// 識別子
		if t.isAlphabet(t.currentRune()) || t.currentRune() == '_' {
			startedAt := t.currentPos.Clone()
			ident := t.consumeIdent()
			cur = t.linkNewIdentToken(cur, startedAt, ident)
			// consumeで移動済み
			continue
		}
		return nil, fmt.Errorf("[%s] unexpected rune: %s", t.currentPos.Clone().String(), string(t.currentRune()))
	}

	cur = t.linkNewEofToken(cur, t.currentPos.Clone())
	return head.Next, nil
}
