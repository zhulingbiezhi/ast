package main

import (
	"go/ast"
	"log"
)

func parseVarObj(pre string, obj *ast.Object) {
	pre += prefix
	log.Printf("%sparseVarObj: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)
	switch t := obj.Decl.(type) {
	case *ast.Field:
		parseType(pre, t.Type)
		//parseField(pre,t)
	}
}

func parseFuncObj(pre string, obj *ast.Object) {
	pre += prefix
	log.Printf("%sparseFuncObj: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)

}

func parsePkgObj(pre string, obj *ast.Object) {
	pre += prefix
	log.Printf("%sparsePkgObj: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)

}
func parseConObj(pre string, obj *ast.Object) {
	pre += prefix
	log.Printf("%sparseConObj: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)

}
func parseTypObj(pre string, obj *ast.Object) {
	pre += prefix
	log.Printf("%sparseTypObj: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)
	switch t := obj.Decl.(type) {
	case *ast.TypeSpec:
		parseType(pre, t.Type)
	}
}
func parseLblObj(pre string, obj *ast.Object) {
	pre += prefix
	log.Printf("%sparseLblObjj: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)

}
