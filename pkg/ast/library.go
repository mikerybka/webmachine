package myast

type Library struct {
	Constants map[string]Constant
	Variables map[string]Variable
	Types     map[string]Type
	Functions map[string]Function
}
