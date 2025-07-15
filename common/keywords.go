package common

type Keyword int

const (
	KClass Keyword = iota
	KConstructor
	KFunction
	KMethod
	KField
	KStatic
	KVar
	KInt
	KChar
	KBoolean
	KVoid
	KTrue
	KFalse
	KNull
	KThis
	KLet
	KDo
	KIf
	KElse
	KWhile
	KReturn
)

var StringToKeyword = map[string]Keyword{
	"class":       KClass,
	"constructor": KConstructor,
	"function":    KFunction,
	"method":      KMethod,
	"field":       KField,
	"static":      KStatic,
	"var":         KVar,
	"int":         KInt,
	"char":        KChar,
	"boolean":     KBoolean,
	"void":        KVoid,
	"true":        KTrue,
	"false":       KFalse,
	"null":        KNull,
	"this":        KThis,
	"let":         KLet,
	"do":          KDo,
	"if":          KIf,
	"else":        KElse,
	"while":       KWhile,
	"return":      KReturn,
}

var KeywordToString = map[Keyword]string{
	KClass:       "class",
	KConstructor: "constructor",
	KFunction:    "function",
	KMethod:      "method",
	KField:       "field",
	KStatic:      "static",
	KVar:         "var",
	KInt:         "int",
	KChar:        "char",
	KBoolean:     "boolean",
	KVoid:        "void",
	KTrue:        "true",
	KFalse:       "false",
	KNull:        "null",
	KThis:        "this",
	KLet:         "let",
	KDo:          "do",
	KIf:          "if",
	KElse:        "else",
	KWhile:       "while",
	KReturn:      "return",
}
