package myast

type ExpressionKind int

const (
	ExpressionKindLiteral ExpressionKind = iota
	ExpressionKindCall
	ExpressionKindRef
)

type Expression struct {
	Kind    ExpressionKind
	Literal *Literal
	Call    *FunctionCall
	Ref     *Reference
}
