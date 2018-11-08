package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

func AstDirParse() {
	fset := token.NewFileSet() // positions are relative to fset
	fs, err := parser.ParseDir(fset, "/Users/huhai/test/golang/src/ast/code",
		func(info os.FileInfo) bool {
			return true
		}, parser.ParseComments)
	if err != nil {
		panic(err)
		return
	}
	for packName, pack := range fs {
		fmt.Printf("package: %s \n", packName)
		parsePackage(pack)
	}
}

func main() {
	AstDirParse()
}
