package tmp

import (
	"fmt"
	"go/ast"
	"reflect"
)

func parseObj(prefix string, obj *ast.Object) {
	fmt.Printf("%s[ast.Object] name: %s kind: %s\n", prefix, obj.Name, obj.Kind.String())
	parseDecl(prefix+"-", obj.Decl)
	parseInterface(prefix+"-", obj.Type)
	parseInterface(prefix+"-", obj.Data)
}

func parseInterface(prefix string, typ interface{}) {
	if typ == nil {
		return
	}
	t := reflect.TypeOf(typ)
	v := reflect.ValueOf(typ)
	v = indirect(v)
	fmt.Printf("type: %s value: %#v\n", t.Name(), v)
}

func parseDecl(prefix string, decl interface{}) {
	switch d := decl.(type) {
	case *ast.Field:
		fmt.Printf("%s[ast.Field] %#v\n", prefix, d)
	case *ast.ImportSpec:
		fmt.Printf("%s[ast.ImportSpec] %#v\n", prefix, d)
	case *ast.ValueSpec:
		fmt.Printf("%s[ast.ValueSpec] %#v\n", prefix, d)
	case *ast.TypeSpec:
		parseTypeSpec(prefix, d)
	case *ast.FuncDecl:
		//fmt.Printf("%s[ast.FuncDecl] %#v\n", prefix, d)
		parseFuncDecl(prefix, d)
	case *ast.LabeledStmt:
		fmt.Printf("%s[ast.LabeledStmt] %#v\n", prefix, d)
	case *ast.AssignStmt:
		fmt.Printf("%s[ast.AssignStmt] %#v\n", prefix, d)
	case *ast.Scope:
		fmt.Printf("%s[ast.Scope] %#v\n", prefix, d)
	default:
		fmt.Printf("%s[unknown decl] %#v\n", prefix, d)
	}
}

func parseTypeSpec(prefix string, t *ast.TypeSpec) {
	switch s := t.Type.(type) {
	case *ast.StructType:
		fmt.Printf("%s[ast.StructType] name: %s- %#v \n", prefix, t.Name, s)
	case *ast.ArrayType:
		fmt.Printf("%s[ast.ArrayType] %#v\n", prefix, s)
	case *ast.ChanType:
		fmt.Printf("%s[ast.ChanType] %#v\n", prefix, s)
	case *ast.FuncType:
		fmt.Printf("%s[ast.FuncType] %#v\n", prefix, s)
	case *ast.InterfaceType:
		fmt.Printf("%s[ast.InterfaceType] %#v\n", prefix, s)
	case *ast.MapType:
		fmt.Printf("%s[ast.MapType] %#v\n", prefix, s)
	default:
		fmt.Printf("%s[unknown typeSpec] %#v\n", prefix, s)
	}
}

func parseFuncDecl(prefix string, decl *ast.FuncDecl) {
	for _, param := range decl.Type.Params.List {
		for _, name := range param.Names {
			fmt.Printf("%s[ast.FuncDecl] name: %s,type: %#v\n", prefix, name.Name, name)
			parseObj(prefix+"-", name.Obj)
		}
	}
}
