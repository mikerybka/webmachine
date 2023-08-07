package myast

type StatementKind int

const (
	StatementKindReturn StatementKind = iota
)

type Statement struct {
	Kind   StatementKind
	Return *ReturnStatement
}

type ReturnStatement struct {
	Values []Expression
}
