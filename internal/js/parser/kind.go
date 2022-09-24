package parser

import (
	commonParser "dragon/pkg/parser"
	commonTokenizer "dragon/pkg/tokenizer"
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
	OR

	Lt
	Le
	Gt
	Ge

	Add
	Sub

	Mul
	Div
	Mod

	ObjectAccess
	Field

	Parenthesis
	Identifier
	Call
	CallArgs
	String
	Number
	Boolean
	Array
	Object
	Null
	NaN
	Undefined
)

var kinds = [...]string{
	FuncDefine:   "FuncDefine",
	FuncParams:   "FuncParams",
	Return:       "Return",
	If:           "If",
	IfElse:       "IfElse",
	While:        "While",
	For:          "For",
	Block:        "Block",
	Assign:       "Assign",
	VarDeclare:   "VarDeclare",
	VarDefine:    "VarDefine",
	LetDeclare:   "LetDeclare",
	LetDefine:    "LetDefine",
	ConstDefine:  "ConstDefine",
	And:          "And",
	OR:           "OR",
	Lt:           "Lt",
	Le:           "Le",
	Gt:           "Gt",
	Ge:           "Ge",
	Add:          "Add",
	Sub:          "Sub",
	Mul:          "Mul",
	Div:          "Div",
	Mod:          "Mod",
	ObjectAccess: "ObjectAccess",
	Field:        "Field",
	Parenthesis:  "Parenthesis",
	Identifier:   "Identifier",
	Call:         "Call",
	CallArgs:     "CallArgs",
	String:       "String",
	Number:       "Number",
	Boolean:      "Boolean",
	Array:        "Array",
	Object:       "Object",
	Null:         "Null",
	NaN:          "NaN",
	Undefined:    "Undefined",
}

func (k Kind) String() string {
	return kinds[k]
}

func NewIdentifierNode(t *commonTokenizer.Token) *commonParser.Node {
	return commonParser.NewNode(Identifier, nil, nil, nil, nil, t.S, 0, 0, false)
}

func NewNodeWithOutImmediate(kind commonParser.NodeKind, lhs, rhs, child *commonParser.Node, children []*commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(kind, lhs, rhs, child, children, "", 0, 0, false)
}
