package tokenizer

import (
	"github.com/google/go-cmp/cmp"
	"github.com/x0y14/dragon/common/tokenizer"
	"testing"
)

func TestTokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expectErr   error
		expectToken *tokenizer.Token
	}{
		{
			"1",
			"<h1>hello, world</h1>",
			nil,
			nil,
		},
		{
			"1",
			"<div p1=1 p2=\"example.com\"><!--comment-->&amp</div>",
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			tokenizer_ := NewTokenizer()
			tok, err := tokenizer_.Tokenize([]rune(tt.in))
			if diff := cmp.Diff(tt.expectErr, err); diff != "" {
				t.Errorf("Tokenize() Err mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.expectToken, tok); diff != "" {
				t.Errorf("Tokenize() Err mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
