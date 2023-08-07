package myast

type FunctionCall struct {
	Function string
	Inputs   []Expression
}
