package main

type astResult struct {
	StructMap map[string]*astStruct
	FuncMap map[string]*astFunc
}

type astFunc struct {
	Args []*Field
	Name string
}

type astStruct struct {
	Fields []*Field
	Name string
	IsPointer bool
}

type Field struct {
	Tag string
	Name string
	Type interface{}
}
