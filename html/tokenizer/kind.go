package tokenizer

type Kind int

const (
	_ Kind = iota
	Illegal
	Eof
	Whitespace

	TagBegin    // <
	TagEnd      // >
	Exclamation // !
	Assign      // =
	Hyphen      // -
	Slash       // /
	Amp         // &

	String
	Integer
	Decimal

	Text
)

var kinds = [...]string{
	Illegal:     "Illegal",
	Eof:         "Eof",
	Whitespace:  "Whitespace",
	TagBegin:    "TagBegin",
	TagEnd:      "TagEnd",
	Exclamation: "Exclamation",
	Assign:      "Assign",
	Hyphen:      "Hyphen",
	Slash:       "Slash",
	Amp:         "Amp",
	String:      "String",
	Integer:     "Integer",
	Decimal:     "Decimal",
	Text:        "Text",
}

func (k Kind) String() string {
	return kinds[k]
}

func symbolKind(symbol string) Kind {
	var kind Kind
	switch symbol {
	case "<":
		kind = TagBegin
	case ">":
		kind = TagEnd
	case "!":
		kind = Exclamation
	case "=":
		kind = Assign
	case "-":
		kind = Hyphen
	case "/":
		kind = Slash
	case "&":
		kind = Amp
	default:
		kind = Illegal
	}
	return kind
}
