package tokenizer

type Kind int

const (
	_ Kind = iota
	Illegal
	Eof

	Identifier
	String  // "...", '..."
	Integer // 123
	Decimal // 123.4

	OPAdd       // +
	OPSub       // -
	OPMul       // *
	OPDiv       // /
	OPMod       // %
	OPAssign    // =
	OPAddAssign // +=
	OPSubAssign // -=
	OPMulAssign // *=
	OPDivAssign // /=
	OPModAssign // %=

	LEq  // ==
	LNe  // !=
	LLt  // <
	LLe  // <=
	LGt  // >
	LGe  // >=
	LAnd // &&
	LOr  // ||
	LNot // !

	SYLcb   // {
	SYRcb   // }
	SYLrb   // (
	SYRrb   // )
	SYLsb   // [
	SYRsb   // ]
	SYColon // :
	SYSemi  // ;
	SYDot   // .
	SYComma // ,
	SYQuest // ?

	KWAbstract
	KWArguments
	KWAwait

	KWBoolean
	KWBreak
	KWByte

	KWCase
	KWCatch
	KWChar
	KWClass
	KWConst
	KWContinue

	KWDebugger
	KWDefault
	KWDelete
	KWDo
	KWDouble

	KWElse
	KWEnum
	KWEval
	KWExport
	KWExtends

	KWFalse
	KWFinal
	KWFinally
	KWFloat
	KWFor
	KWFunction

	KWGoto

	KWIf
	KWImplements
	KWImport
	KWIn
	KWInstanceof
	KWInt
	KWInterface

	KWLet
	KWLong

	KWNative
	KWNew
	KWNull

	KWPackage
	KWPrivate
	KWProtected
	KWPublic
	KWReturn
	KWShort
	KWStatic
	KWSuper
	KWSwitch
	KWSynchronized

	KWThis
	KWThrow
	KWThrows
	KWTransient
	KWTrue
	KWTry
	KWTypeof

	KWVar
	KWVoid
	KWVolatile

	KWWhile
	KWWith

	KWYield
)

var kinds = [...]string{
	Eof:            "Eof",
	Illegal:        "Illegal",
	Identifier:     "Identifier",
	String:         "String",
	Integer:        "Integer",
	Decimal:        "Decimal",
	OPAdd:          "OPAdd",
	OPSub:          "OPSub",
	OPMul:          "OPMul",
	OPDiv:          "OPDiv",
	OPMod:          "OPMod",
	OPAssign:       "OPAssign",
	OPAddAssign:    "OPAddAssign",
	OPSubAssign:    "OPSubAssign",
	OPMulAssign:    "OPMulAssign",
	OPDivAssign:    "OPDivAssign",
	OPModAssign:    "OPModAssign",
	LEq:            "LEq",
	LNe:            "LNe",
	LLt:            "LLt",
	LLe:            "LLe",
	LGt:            "LGt",
	LGe:            "LGe",
	LAnd:           "LAnd",
	LOr:            "LOr",
	LNot:           "LNot",
	SYLcb:          "SYLcb",
	SYRcb:          "SYRcb",
	SYLrb:          "SYLrb",
	SYRrb:          "SYRrb",
	SYLsb:          "SYLsb",
	SYRsb:          "SYRsb",
	SYColon:        "SYColon",
	SYSemi:         "SYSemi",
	SYDot:          "SYDot",
	SYComma:        "SYComma",
	SYQuest:        "SYQuest",
	KWAbstract:     "KWAbstract",
	KWArguments:    "KWArguments",
	KWAwait:        "KWAwait",
	KWBoolean:      "KWBoolean",
	KWBreak:        "KWBreak",
	KWByte:         "KWByte",
	KWCase:         "KWCase",
	KWCatch:        "KWCatch",
	KWChar:         "KWChar",
	KWClass:        "KWClass",
	KWConst:        "KWConst",
	KWContinue:     "KWContinue",
	KWDebugger:     "KWDebugger",
	KWDefault:      "KWDefault",
	KWDelete:       "KWDelete",
	KWDo:           "KWDo",
	KWDouble:       "KWDouble",
	KWElse:         "KWElse",
	KWEnum:         "KWEnum",
	KWEval:         "KWEval",
	KWExport:       "KWExport",
	KWExtends:      "KWExtends",
	KWFalse:        "KWFalse",
	KWFinal:        "KWFinal",
	KWFinally:      "KWFinally",
	KWFloat:        "KWFloat",
	KWFor:          "KWFor",
	KWFunction:     "KWFunction",
	KWGoto:         "KWGoto",
	KWIf:           "KWIf",
	KWImplements:   "KWImplements",
	KWImport:       "KWImport",
	KWIn:           "KWIn",
	KWInstanceof:   "KWInstanceof",
	KWInt:          "KWInt",
	KWInterface:    "KWInterface",
	KWLet:          "KWLet",
	KWLong:         "KWLong",
	KWNative:       "KWNative",
	KWNew:          "KWNew",
	KWNull:         "KWNull",
	KWPackage:      "KWPackage",
	KWPrivate:      "KWPrivate",
	KWProtected:    "KWProtected",
	KWPublic:       "KWPublic",
	KWReturn:       "KWReturn",
	KWShort:        "KWShort",
	KWStatic:       "KWStatic",
	KWSuper:        "KWSuper",
	KWSwitch:       "KWSwitch",
	KWSynchronized: "KWSynchronized",
	KWThis:         "KWThis",
	KWThrow:        "KWThrow",
	KWThrows:       "KWThrows",
	KWTransient:    "KWTransient",
	KWTrue:         "KWTrue",
	KWTry:          "KWTry",
	KWTypeof:       "KWTypeof",
	KWVar:          "KWVar",
	KWVoid:         "KWVoid",
	KWVolatile:     "KWVolatile",
	KWWhile:        "KWWhile",
	KWWith:         "KWWith",
	KWYield:        "KWYield",
}

func (k Kind) String() string {
	return kinds[k]
}

func SymbolKind(symbol string) Kind {
	switch symbol {
	case "+":
		return OPAdd
	case "-":
		return OPSub
	case "*":
		return OPMul
	case "/":
		return OPDiv
	case "%":
		return OPMod
	case "=":
		return OPAssign
	case "+=":
		return OPAddAssign
	case "-=":
		return OPSubAssign
	case "*=":
		return OPMulAssign
	case "/=":
		return OPDivAssign
	case "%=":
		return OPModAssign
	case "==":
		return LEq
	case "!=":
		return LNe
	case "<":
		return LLt
	case "<=":
		return LLe
	case ">":
		return LGt
	case ">=":
		return LGe
	case "&&":
		return LAnd
	case "||":
		return LOr
	case "!":
		return LNot
	case "{":
		return SYLcb
	case "}":
		return SYRcb
	case "(":
		return SYLrb
	case ")":
		return SYRrb
	case "[":
		return SYLsb
	case "]":
		return SYRsb
	case ":":
		return SYColon
	case ";":
		return SYSemi
	case ".":
		return SYDot
	case ",":
		return SYComma
	case "?":
		return SYQuest
	}
	return Illegal
}

func IdentKind(ident string) Kind {
	switch ident {
	case "abstract":
		return KWAbstract
	case "arguments":
		return KWArguments
	case "await":
		return KWAwait
	case "boolean":
		return KWBoolean
	case "break":
		return KWBreak
	case "byte":
		return KWByte
	case "case":
		return KWCase
	case "catch":
		return KWCatch
	case "char":
		return KWChar
	case "class":
		return KWClass
	case "const":
		return KWConst
	case "continue":
		return KWContinue
	case "debugger":
		return KWDebugger
	case "default":
		return KWDefault
	case "delete":
		return KWDelete
	case "do":
		return KWDo
	case "double":
		return KWDouble
	case "else":
		return KWElse
	case "enum":
		return KWEnum
	case "eval":
		return KWEval
	case "export":
		return KWExport
	case "extends":
		return KWExtends
	case "false":
		return KWFalse
	case "final":
		return KWFinal
	case "finally":
		return KWFinally
	case "float":
		return KWFloat
	case "for":
		return KWFor
	case "function":
		return KWFunction
	case "goto":
		return KWGoto
	case "if":
		return KWIf
	case "implements":
		return KWImplements
	case "import":
		return KWImport
	case "in":
		return KWIn
	case "instanceof":
		return KWInstanceof
	case "int":
		return KWInt
	case "interface":
		return KWInterface
	case "let":
		return KWLet
	case "long":
		return KWLong
	case "native":
		return KWNative
	case "new":
		return KWNew
	case "null":
		return KWNull
	case "package":
		return KWPackage
	case "private":
		return KWPrivate
	case "protected":
		return KWProtected
	case "public":
		return KWPublic
	case "return":
		return KWReturn
	case "short":
		return KWShort
	case "static":
		return KWStatic
	case "super":
		return KWSuper
	case "switch":
		return KWSwitch
	case "synchronized":
		return KWSynchronized
	case "this":
		return KWThis
	case "throw":
		return KWThrow
	case "throws":
		return KWThrows
	case "transient":
		return KWTransient
	case "true":
		return KWTrue
	case "try":
		return KWTry
	case "typeof":
		return KWTypeof
	case "var":
		return KWVar
	case "void":
		return KWVoid
	case "volatile":
		return KWVolatile
	case "while":
		return KWWhile
	case "with":
		return KWWith
	case "yield":
		return KWYield
	}
	return Identifier
}
