package tokenizer

type Kind int

const (
	_ Kind = iota

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

func (k Kind) String() string {
	return ""
}
