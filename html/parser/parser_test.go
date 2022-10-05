package parser

import (
	"github.com/google/go-cmp/cmp"
	commonParser "github.com/x0y14/dragon/common/parser"
	"github.com/x0y14/dragon/html/tokenizer"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		in            string
		expectedErr   error
		expectedNodes []*commonParser.Node
	}{
		{
			"comment",
			"<!--hello-->",
			nil,
			[]*commonParser.Node{
				commonParser.NewNode(Comment,
					nil, nil, nil,
					NewImmediateNode(String, "hello", 0, 0, false), nil, nil,
					"", 0, 0, false),
			},
		},
		{
			"comment",
			"<!--  hello, world  -->",
			nil,
			[]*commonParser.Node{
				commonParser.NewNode(Comment,
					nil, nil, nil,
					NewImmediateNode(String, " hello, world ", 0, 0, false), nil, nil,
					"", 0, 0, false),
			},
		},
		{
			"doctype",
			"<!doctype html>",
			nil,
			[]*commonParser.Node{
				commonParser.NewNode(Doctype,
					nil, nil, nil,
					NewImmediateNode(String, "html", 0, 0, false), nil, nil,
					"", 0, 0, false),
			},
		},
		{
			"h1",
			"<h1>hello</h1>",
			nil,
			[]*commonParser.Node{
				NewTagNode(NewIdentifierNode("h1"), nil, []*commonParser.Node{
					NewTextNode("hello"),
				}),
			},
		},
		{
			"h1 br",
			"<h1>hello,<br /> world</h1>",
			nil,
			[]*commonParser.Node{
				NewTagNode(NewIdentifierNode("h1"), nil, []*commonParser.Node{
					NewTextNode("hello,"),
					NewSoloTagNode(NewIdentifierNode("br"),
						nil,
					),
					NewTextNode(" world"),
				}),
			},
		},
		{
			"h1 br",
			"<h1 style=\"text-align:right\">hello,<img src=\"google.com\" /> world</h1>",
			nil,
			[]*commonParser.Node{
				NewTagNode(NewIdentifierNode("h1"), NewParametersNode([]*commonParser.Node{
					NewParamNode(NewIdentifierNode("style"), NewStringNode("text-align:right")),
				}), []*commonParser.Node{
					NewTextNode("hello,"),
					NewSoloTagNode(NewIdentifierNode("img"),
						NewParametersNode([]*commonParser.Node{
							NewParamNode(NewIdentifierNode("src"), NewStringNode("google.com")),
						}),
					),
					NewTextNode(" world"),
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer_ := tokenizer.NewTokenizer()
			tok, err := tokenizer_.Tokenize([]rune(tt.in))
			if err != nil {
				t.Fatalf("failed to tokenize: %v", err)
			}
			parser_ := NewParser()
			nodes, err := parser_.Parse(tok)

			if diff := cmp.Diff(tt.expectedErr, err); diff != "" {
				t.Errorf("Parse() Err mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedNodes, nodes); diff != "" {
				t.Errorf("Parse() Err mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
