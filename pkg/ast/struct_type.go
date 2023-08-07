package myast

import "go/ast"

type StructType struct {
	Fields []*Field `json:"fields"`
}

func (s *StructType) Visit(n ast.Node) *Type {
	t := &Type{
		Kind:   StructTypeKind,
		Struct: s,
	}
	switch n := n.(type) {
	case *ast.StructType:
		for _, field := range n.Fields.List {
			s.Fields = append(s.Fields, (&Field{}).Visit(field))
		}
		return t
	default:
		panic("unreachable")
	}
}
