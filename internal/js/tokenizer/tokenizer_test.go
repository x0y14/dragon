package tokenizer

import (
	"dragon/pkg/tokenizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expectErr   error
		expectToken *tokenizer.Token
	}{
		{ // 出典: https://www.programiz.com/javascript/examples/add-number
			"add 2 nums",
			`
				const num1 = 5;
				const num2 = 3;
				
				// add two numbers
				const sum = num1 + num2;
				
				// display the sum
				console.log('The sum of ' + num1 + ' and ' + num2 + ' is: ' + sum);`,
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer_ := NewTokenizer()
			tok, err := tokenizer_.Tokenize([]rune(tt.in))
			if !assert.Equal(t, tt.expectErr, err) {
				t.Fatalf("%v", err)
			}
			assert.Equal(t, tt.expectToken, tok)
		})
	}
}
