package parser

import (
	"dragon/common/parser"
	"dragon/js/tokenizer"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expectErr   error
		expectNodes []*parser.Node
	}{
		{
			"hello",
			"console.log(\"hello\")",
			nil,
			[]*parser.Node{
				{
					Kind: Field,
					LHS:  &parser.Node{Kind: Identifier, S: "console"},
					RHS: &parser.Node{Kind: Call,
						N1: &parser.Node{Kind: Identifier, S: "log"},
						N2: &parser.Node{
							Kind: CallArgs,
							Children: []*parser.Node{
								{Kind: String, S: "hello"},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer_ := tokenizer.NewTokenizer()
			token_, err := tokenizer_.Tokenize([]rune(tt.in))
			if err != nil {
				t.Fatalf("failed to tokenize: %v", err)
			}
			parser_ := NewParser()
			nodes, err := parser_.Parse(token_)
			if diff := cmp.Diff(tt.expectErr, err); diff != "" {
				t.Errorf("Parse() Err mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.expectNodes, nodes); diff != "" {
				t.Errorf("Parse() Nodes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
