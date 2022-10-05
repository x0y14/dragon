package parser

import (
	commonParser "github.com/x0y14/dragon/common/parser"
)

type Kind int

const (
	_ Kind = iota
	Parameters
	Parameter
	Identifier
	String
	Text

	Tag
	SoloTag

	Comment
	Doctype
)

var kinds = [...]string{
	Parameters: "Parameters",
	Parameter:  "Parameter",
	Identifier: "Identifier",
	String:     "String",
	Text:       "Text",
	Tag:        "Tag",
	SoloTag:    "SoloTag",
	Comment:    "Comment",
	Doctype:    "Doctype",
}

func (k Kind) String() string {
	return kinds[k]
}

func NewImmediateNode(k Kind, s string, f float64, i int64, b bool) *commonParser.Node {
	return commonParser.NewNode(k, nil, nil, nil, nil, nil, nil, s, f, i, b)
}

func NewStringNode(s string) *commonParser.Node {
	return commonParser.NewNode(String, nil, nil, nil, nil, nil, nil, s, 0, 0, false)
}

func NewTextNode(s string) *commonParser.Node {
	return commonParser.NewNode(Text, nil, nil, nil, nil, nil, nil, s, 0, 0, false)
}

func NewIdentifierNode(s string) *commonParser.Node {
	return commonParser.NewNode(Identifier, nil, nil, nil, nil, nil, nil, s, 0, 0, false)
}

func NewParamNode(ident, str *commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(Parameter, ident, str, nil, nil, nil, nil, "", 0, 0, false)
}

func NewParametersNode(params []*commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(Parameters, nil, nil, params, nil, nil, nil, "", 0, 0, false)
}

func NewSoloTagNode(ident, params *commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(SoloTag, nil, nil, nil, ident, params, nil, "", 0, 0, false)
}

func NewTagNode(ident, params *commonParser.Node, children []*commonParser.Node) *commonParser.Node {
	return commonParser.NewNode(Tag, nil, nil, children, ident, params, nil, "", 0, 0, false)
}
