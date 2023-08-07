package myast

import (
	"go/ast"
	"reflect"
)

type TypeKind string

const (
	PrimitiveTypeKind = TypeKind("primitive")
	PointerTypeKind   = TypeKind("pointer")
	StructTypeKind    = TypeKind("struct")
	MapTypeKind       = TypeKind("map")
	ArrayTypeKind     = TypeKind("array")
	InterfaceTypeKind = TypeKind("interface")
)

type Type struct {
	Kind      TypeKind    `json:"kind"`
	Primitive string      `json:"primitive,omitempty"`
	Pointer   *Type       `json:"pointer,omitempty"`
	Struct    *StructType `json:"struct,omitempty"`
	ValueType *Type       `json:"valueType,omitempty"`
	Methods   []Method    `json:"methods,omitempty"`
}

var types = map[string]*Type{}

func TypeFromReflect(t reflect.Type) *Type {
	typ, ok := types[t.String()]
	if ok {
		return typ
	} else {
		typ = &Type{}
		types[t.String()] = typ
	}
	switch t.Kind() {
	case reflect.Invalid:
		panic("invalid type")
	case reflect.Bool:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "bool"
	case reflect.Int:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "int"
	case reflect.Int8:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "int8"
	case reflect.Int16:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "int16"
	case reflect.Int32:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "int32"
	case reflect.Int64:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "int64"
	case reflect.Uint:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "uint"
	case reflect.Uint8:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "uint8"
	case reflect.Uint16:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "uint16"
	case reflect.Uint32:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "uint32"
	case reflect.Uint64:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "uint64"
	case reflect.Uintptr:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "uintptr"
	case reflect.Float32:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "float32"
	case reflect.Float64:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "float64"
	case reflect.Complex64:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "complex64"
	case reflect.Complex128:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "complex128"
	case reflect.Array:
		typ.Kind = ArrayTypeKind
		typ.ValueType = TypeFromReflect(t.Elem())
	case reflect.Chan:
		panic("channels not supported")
	case reflect.Func:
		panic("functions not supported")
	case reflect.Interface:
		typ.Kind = InterfaceTypeKind
	case reflect.Map:
		typ.Kind = MapTypeKind
		typ.ValueType = TypeFromReflect(t.Elem())
	case reflect.Pointer:
		typ = TypeFromReflect(t.Elem())
	case reflect.Slice:
		typ.Kind = ArrayTypeKind
		typ.ValueType = TypeFromReflect(t.Elem())
	case reflect.String:
		typ.Kind = PrimitiveTypeKind
		typ.Primitive = "string"
	case reflect.Struct:
		typ.Kind = StructTypeKind
		typ.Struct = &StructType{}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			typ.Struct.Fields = append(typ.Struct.Fields, &Field{
				Name: f.Name,
				Type: f.Type.String(),
			})
		}
	case reflect.UnsafePointer:
		panic("unsafe pointers not supported")
	default:
		panic("unreachable")
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		meth := Method{
			Name: m.Name,
		}
		for i := 1; i < m.Type.NumIn(); i++ {
			meth.Params = append(meth.Params, &Field{
				Name: m.Type.In(i).Name(),
				Type: m.Type.In(i).String(),
			})
		}
		for i := 0; i < m.Type.NumOut(); i++ {
			output := m.Type.Out(i)
			meth.Results = append(meth.Results, &Field{
				Name: output.Name(),
				Type: output.String(),
			})
		}
		typ.Methods = append(typ.Methods, meth)
	}

	return typ
}

func TypeOf(v any) *Type {
	t := reflect.TypeOf(v)
	return TypeFromReflect(t)
}

func (t *Type) Visit(n ast.Node) *Type {
	switch n := n.(type) {
	case *ast.Ident:
		t.Kind = PrimitiveTypeKind
		t.Primitive = n.Name
		return t
	case *ast.StarExpr:
		t.Kind = PointerTypeKind
		t.Pointer = &Type{}
		return t.Pointer.Visit(n.X)
	case *ast.StructType:
		t.Kind = StructTypeKind
		t.Struct = &StructType{}
		return t.Struct.Visit(n)
	case *ast.MapType:
		t.Kind = MapTypeKind
		return t.ValueType.Visit(n.Value)
	case *ast.ArrayType:
		t.Kind = ArrayTypeKind
		return t.ValueType.Visit(n.Elt)
	default:
		panic("unreachable")
	}
}

var (
	StringType = &Type{
		Kind:      PrimitiveTypeKind,
		Primitive: "string",
	}
)
