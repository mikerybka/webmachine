package myast

import "go/ast"

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (f *Field) Visit(n ast.Node) *Field {
	switch n := n.(type) {
	case *ast.Field:
		f.Name = n.Names[0].Name
		f.Type = n.Type.(*ast.Ident).Name
		return f
	default:
		panic("unreachable")
	}
}
