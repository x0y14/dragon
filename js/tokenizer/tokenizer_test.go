package tokenizer

import (
	"dragon/common/tokenizer"
	"github.com/google/go-cmp/cmp"
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
			&tokenizer.Token{
				Kind: KWConst,
				Pos:  tokenizer.NewPositionFromString("\n"),
				S:    "const",
				I:    0,
				F:    0,
				Next: &tokenizer.Token{
					Kind: Identifier,
					Pos:  tokenizer.NewPositionFromString("\nconst "),
					S:    "num1",
					I:    0,
					F:    0,
					Next: &tokenizer.Token{
						Kind: OPAssign,
						Pos:  tokenizer.NewPositionFromString("\nconst num1 "),
						S:    "",
						I:    0,
						F:    0,
						Next: &tokenizer.Token{
							Kind: Integer,
							Pos:  tokenizer.NewPositionFromString("\nconst num1 = "),
							S:    "",
							I:    5,
							F:    0,
							Next: &tokenizer.Token{
								Kind: SYSemi,
								Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5"),
								S:    "",
								I:    0,
								F:    0,
								Next: &tokenizer.Token{
									Kind: KWConst,
									Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\n"),
									S:    "const",
									I:    0,
									F:    0,
									Next: &tokenizer.Token{
										Kind: Identifier,
										Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst "),
										S:    "num2",
										I:    0,
										F:    0,
										Next: &tokenizer.Token{
											Kind: OPAssign,
											Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 "),
											S:    "",
											I:    0,
											F:    0,
											Next: &tokenizer.Token{
												Kind: Integer,
												Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = "),
												S:    "",
												I:    3,
												F:    0,
												Next: &tokenizer.Token{
													Kind: SYSemi,
													Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3"),
													S:    "",
													I:    0,
													F:    0,
													Next: &tokenizer.Token{
														Kind: KWConst,
														Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\n"),
														S:    "const",
														I:    0,
														F:    0,
														Next: &tokenizer.Token{
															Kind: Identifier,
															Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst "),
															S:    "sum",
															I:    0,
															F:    0,
															Next: &tokenizer.Token{
																Kind: OPAssign,
																Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum "),
																S:    "",
																I:    0,
																F:    0,
																Next: &tokenizer.Token{
																	Kind: Identifier,
																	Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = "),
																	S:    "num1",
																	I:    0,
																	F:    0,
																	Next: &tokenizer.Token{
																		Kind: OPAdd,
																		Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 "),
																		S:    "",
																		I:    0,
																		F:    0,
																		Next: &tokenizer.Token{
																			Kind: Identifier,
																			Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + "),
																			S:    "num2",
																			I:    0,
																			F:    0,
																			Next: &tokenizer.Token{
																				Kind: SYSemi,
																				Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2"),
																				S:    "",
																				I:    0,
																				F:    0,
																				Next: &tokenizer.Token{
																					Kind: Identifier,
																					Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\n"),
																					S:    "console",
																					I:    0,
																					F:    0,
																					Next: &tokenizer.Token{
																						Kind: SYDot,
																						Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole"),
																						S:    "",
																						I:    0,
																						F:    0,
																						Next: &tokenizer.Token{
																							Kind: Identifier,
																							Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole."),
																							S:    "log",
																							I:    0,
																							F:    0,
																							Next: &tokenizer.Token{
																								Kind: SYLrb,
																								Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log"),
																								S:    "",
																								I:    0,
																								F:    0,
																								Next: &tokenizer.Token{
																									Kind: String,
																									Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log("),
																									S:    "The sum of ",
																									I:    0,
																									F:    0,
																									Next: &tokenizer.Token{
																										Kind: OPAdd,
																										Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' "),
																										S:    "",
																										I:    0,
																										F:    0,
																										Next: &tokenizer.Token{
																											Kind: Identifier,
																											Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + "),
																											S:    "num1",
																											I:    0,
																											F:    0,
																											Next: &tokenizer.Token{
																												Kind: OPAdd,
																												Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 "),
																												S:    "",
																												I:    0,
																												F:    0,
																												Next: &tokenizer.Token{
																													Kind: String,
																													Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + "),
																													S:    " and ",
																													I:    0,
																													F:    0,
																													Next: &tokenizer.Token{
																														Kind: OPAdd,
																														Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' "),
																														S:    "",
																														I:    0,
																														F:    0,
																														Next: &tokenizer.Token{
																															Kind: Identifier,
																															Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + "),
																															S:    "num2",
																															I:    0,
																															F:    0,
																															Next: &tokenizer.Token{
																																Kind: OPAdd,
																																Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 "),
																																S:    "",
																																I:    0,
																																F:    0,
																																Next: &tokenizer.Token{
																																	Kind: String,
																																	Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 + "),
																																	S:    " is: ",
																																	I:    0,
																																	F:    0,
																																	Next: &tokenizer.Token{
																																		Kind: OPAdd,
																																		Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 + ' is: ' "),
																																		S:    "",
																																		I:    0,
																																		F:    0,
																																		Next: &tokenizer.Token{
																																			Kind: Identifier,
																																			Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 + ' is: ' + "),
																																			S:    "sum",
																																			I:    0,
																																			F:    0,
																																			Next: &tokenizer.Token{
																																				Kind: SYRrb,
																																				Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 + ' is: ' + sum"),
																																				S:    "",
																																				I:    0,
																																				F:    0,
																																				Next: &tokenizer.Token{
																																					Kind: SYSemi,
																																					Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 + ' is: ' + sum)"),
																																					S:    "",
																																					I:    0,
																																					F:    0,
																																					Next: &tokenizer.Token{
																																						Kind: Eof,
																																						Pos:  tokenizer.NewPositionFromString("\nconst num1 = 5;\nconst num2 = 3;\n\n// add two numbers\nconst sum = num1 + num2;\n\n// display the sum\nconsole.log('The sum of ' + num1 + ' and ' + num2 + ' is: ' + sum);"),
																																						S:    "",
																																						I:    0,
																																						F:    0,
																																						Next: nil,
																																					},
																																				},
																																			},
																																		},
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer_ := NewTokenizer()
			tok, err := tokenizer_.Tokenize([]rune(tt.in))
			if diff := cmp.Diff(tt.expectErr, err); diff != "" {
				t.Errorf("Tokenize() Err mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectToken, tok); diff != "" {
				t.Errorf("Tokenize() Token mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
