package myast_test

import (
	"testing"
)

func TestTypeResolvePath(t *testing.T) {
	// t.Parallel()
	// typ := ast.Type{
	// 	Kind: ast.StructTypeKind,
	// 	Struct: &ast.StructType{
	// 		Fields: []*ast.Field{
	// 			{
	// 				Name: "foo",
	// 				Type: ast.Type{
	// 					Kind:      ast.PrimitiveTypeKind,
	// 					Primitive: "string",
	// 				},
	// 			},
	// 			{
	// 				Name: "bar",
	// 				Type: ast.Type{
	// 					Kind: ast.StructTypeKind,
	// 					Struct: &ast.StructType{
	// 						Fields: []*ast.Field{
	// 							{
	// 								Name: "baz",
	// 								Type: ast.Type{
	// 									Kind:      ast.PrimitiveTypeKind,
	// 									Primitive: "string",
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// foo, ok := typ.ResolvePath("/foo")
	// if !ok {
	// 	t.Fatal("expected to resolve /foo")
	// }
	// if foo.Kind != ast.PrimitiveTypeKind {
	// 	t.Fatal("expected to resolve to primitive type")
	// }
	// if foo.Primitive != "string" {
	// 	t.Fatal("expected to resolve to string")
	// }
	// bar, ok := typ.ResolvePath("/bar")
	// if !ok {
	// 	t.Fatal("expected to resolve /bar")
	// }
	// if bar.Kind != ast.StructTypeKind {
	// 	t.Fatal("expected to resolve to Struct type")
	// }
	// if bar.Struct == nil {
	// 	t.Fatal("expected to resolve to Struct type")
	// }
	// if len(bar.Struct.Fields) != 1 {
	// 	t.Fatal("expected to resolve to Struct type with one field")
	// }
	// if bar.Struct.Fields[0].Name != "baz" {
	// 	t.Fatal("expected to resolve to Struct type with field named baz")
	// }
	// if bar.Struct.Fields[0].Type.Kind != ast.PrimitiveTypeKind {
	// 	t.Fatal("expected to resolve to Struct type with field of primitive type")
	// }
	// if bar.Struct.Fields[0].Type.Primitive != "string" {
	// 	t.Fatal("expected to resolve to Struct type with field of string type")
	// }
	// baz, ok := typ.ResolvePath("/bar/baz")
	// if !ok {
	// 	t.Fatal("expected to resolve /bar/baz")
	// }
	// if baz.Kind != ast.PrimitiveTypeKind {
	// 	t.Fatal("expected to resolve to primitive type")
	// }
	// if baz.Primitive != "string" {
	// 	t.Fatal("expected to resolve to string")
	// }
	// _, ok = typ.ResolvePath("/bar/baz/qux")
	// if ok {
	// 	t.Fatal("expected to not resolve /bar/baz/qux")
	// }
}
