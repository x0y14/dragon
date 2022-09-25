package parser

import (
	commonParser "dragon/common/parser"
	commonTokenizer "dragon/common/tokenizer"
	"dragon/js/tokenizer"
	"fmt"
)

type Kind int

const (
	_ Kind = iota

	FuncDefine
	FuncParams
	Return
	If
	IfElse
	While
	For
	Block

	Assign      // lhs = rhs;
	VarDeclare  // var ident;
	VarDefine   // var lhs = rhs;
	LetDeclare  // let ident;
	LetDefine   // let ident = ...;
	ConstDefine // const ident = ...;

	And
	Or

	Eq
	Ne

	Lt
	Le
	Gt
	Ge

	Add
	Sub

	Mul
	Div
	Mod

	Minus
	Not

	ObjectAccess
	Field

	Parenthesis
	Identifier
	Call
	CallArgs
	String
	NumberDecimal
	NumberInteger
	Boolean
	Array
	Object
	Null
	NaN
	Undefined
)

var kinds = [...]string{
	FuncDefine:    "FuncDefine",
	FuncParams:    "FuncParams",
	Return:        "Return",
	If:            "If",
	IfElse:        "IfElse",
	While:         "While",
	For:           "For",
	Block:         "Block",
	Assign:        "Assign",
	VarDeclare:    "VarDeclare",
	VarDefine:     "VarDefine",
	LetDeclare:    "LetDeclare",
	LetDefine:     "LetDefine",
	ConstDefine:   "ConstDefine",
	And:           "And",
	Or:            "Or",
	Eq:            "Eq",
	Ne:            "Ne",
	Lt:            "Lt",
	Le:            "Le",
	Gt:            "Gt",
	Ge:            "Ge",
	Add:           "Add",
	Sub:           "Sub",
	Mul:           "Mul",
	Div:           "Div",
	Mod:           "Mod",
	Minus:         "Minus",
	Not:           "Not",
	ObjectAccess:  "ObjectAccess",
	Field:         "Field",
	Parenthesis:   "Parenthesis",
	Identifier:    "Identifier",
	Call:          "Call",
	CallArgs:      "CallArgs",
	String:        "String",
	NumberDecimal: "NumberDecimal",
	NumberInteger: "NumberInteger",
	Boolean:       "Boolean",
	Array:         "Array",
	Object:        "Object",
	Null:          "Null",
	NaN:           "NaN",
	Undefined:     "Undefined",
}

func (k Kind) String() string {
	return kinds[k]
}

func NewIdentifierNode(t *commonTokenizer.Token) *commonParser.Node {
	return commonParser.NewNode(Identifier, nil, nil, nil, nil, nil, nil, t.S, 0, 0, false)
}

func NewImmediateNode(t *commonTokenizer.Token) (*commonParser.Node, error) {
	switch t.Kind {
	case tokenizer.String:
		return commonParser.NewNode(String, nil, nil, nil, nil, nil, nil, t.S, 0, 0, false), nil
	case tokenizer.Decimal:
		return commonParser.NewNode(NumberDecimal, nil, nil, nil, nil, nil, nil, "", t.F, 0, false), nil
	case tokenizer.Integer:
		return commonParser.NewNode(NumberInteger, nil, nil, nil, nil, nil, nil, "", 0, t.I, false), nil
	case tokenizer.KWTrue:
		return commonParser.NewNode(Boolean, nil, nil, nil, nil, nil, nil, "", 0, 0, true), nil
	case tokenizer.KWFalse:
		return commonParser.NewNode(Boolean, nil, nil, nil, nil, nil, nil, "", 0, 0, false), nil
	case tokenizer.KWNull:
		return commonParser.NewNode(Null, nil, nil, nil, nil, nil, nil, "", 0, 0, false), nil
	}
	return nil, fmt.Errorf("unsupported token for immediate node: %v", t.Kind.String())
}

func NewNodeWithOutImmediate(kind commonParser.NodeKind, lhs, rhs *commonParser.Node, children []*commonParser.Node, n1, n2, n3 *commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(kind, lhs, rhs, children, n1, n2, n3, "", 0, 0, false)
}

func NewNodeWithOutBothSidesAndImmediate(kind commonParser.NodeKind, children []*commonParser.Node, n1, n2, n3 *commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(kind, nil, nil, children, n1, n2, n3, "", 0, 0, false)
}

func NewNodeBothSides(kind commonParser.NodeKind, lhs, rhs *commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(kind, lhs, rhs, nil, nil, nil, nil, "", 0, 0, false)
}
