package main

import (
	"fmt"
	"go/ast"
	"log"
)

var prefix = "-"

func parsePackage(pkg *ast.Package) {
	parseScope("", pkg.Scope)
	for _, value := range pkg.Files {
		parseFile("", value)
	}
	for _, value := range pkg.Imports {
		parseObject("", value)
	}
}
func parseFile(pre string, file *ast.File) {
	if file == nil {
		return
	}
	pre += prefix
	parseObject(pre, file.Name.Obj)
	for _, value := range file.Unresolved {
		parseObject(pre, value.Obj)
	}
	parseScope(pre, file.Scope)
	for _, value := range file.Decls {
		parseDecl(pre, &value)
	}
}

func parseDecl(pre string, decl *ast.Decl) {
	if decl == nil {
		return
	}
	pre += prefix
	log.Printf("%sparseDecl: %#v\n", pre, *decl)
	//switch d := decl.(type) {
	//case *ast.Field:
	//case *ast.ImportSpec:
	//case *ast.ValueSpec:
	//case *ast.TypeSpec:
	//case *ast.FuncDecl:
	//case *ast.Scope:
	//}
}

func parseScope(pre string, scope *ast.Scope) {
	if scope == nil {
		return
	}
	pre += prefix
	log.Printf("%sparseScope: \n", pre)
	parseScope(pre, scope.Outer)
	for _, value := range scope.Objects {
		parseObject(pre, value)
	}
}

func parseObject(pre string, obj *ast.Object) {
	if obj == nil {
		return
	}
	pre += prefix
	switch obj.Kind {
	case ast.Pkg: // package
		parsePkgObj(pre, obj)
	case ast.Con: // constant
		parseConObj(pre, obj)
	case ast.Typ: // type
		parseTypObj(pre, obj)
	case ast.Var: // variable
		parseVarObj(pre, obj)
	case ast.Fun: // function or method
		parseFuncObj(pre, obj)
	case ast.Lbl: // label
		parseLblObj(pre, obj)
	default:
		log.Printf("%sunknow Object: name:%s kind:%s %#v\n", pre, obj.Name, obj.Kind, obj)
	}
}

func parseType(pre string, typ interface{}) {
	if typ == nil {
		return
	}
	pre += prefix
	log.Printf("%sparseType: %#v\n", pre, typ)
	switch t := typ.(type) {
	case *ast.ArrayType:
	case *ast.ChanType:
	case *ast.MapType:
	case *ast.FuncType:

	case *ast.InterfaceType:
		parseInterface(pre, t)
	case *ast.StructType:
		parseStruct(pre, t)
	case *ast.Ident:
		parseObject(pre, t.Obj)
	default:
		fmt.Printf("%sunknow type: %#v\n", pre, t)
	}
}

func parseInterface(pre string, inter *ast.InterfaceType) {
	log.Printf("%sparseInterface: %#v\n", pre, inter)
	for _, value := range inter.Methods.List {
		parseField(pre, value)
	}
}

func parseStruct(pre string, st *ast.StructType) {
	if st == nil {
		return
	}
	pre += prefix
	log.Printf("%sparseStruct: \n", pre)
	for _, value := range st.Fields.List {
		parseField(pre, value)
	}
}

func parseField(pre string, fd *ast.Field) {
	if fd == nil {
		return
	}
	pre += prefix
	log.Printf("%sparseField: \n", pre)
	for _, value := range fd.Names {
		parseObject(pre, value.Obj)
	}
	if fd.Tag != nil {
		log.Printf("%sparseField.Tag: %s\n", pre, fd.Tag.Value)
	}
	parseType(pre, fd.Type)
}

//
//// parse as enum, in the package, find out all consts with the same type
//func parseIdent(st *ast.Ident, k string, astPkgs []*ast.Package) {
//	m.Title = k
//	basicType := fmt.Sprint(st)
//	if object, isStdLibObject := stdlibObject[basicType]; isStdLibObject {
//		basicType = object
//	}
//	if k, ok := basicTypes[basicType]; ok {
//		typeFormat := strings.Split(k, ":")
//		m.Type = typeFormat[0]
//		m.Format = typeFormat[1]
//	}
//	enums := make(map[int]string)
//	enumValues := make(map[int]interface{})
//	for _, pkg := range astPkgs {
//		for _, fl := range pkg.Files {
//			for _, obj := range fl.Scope.Objects {
//				if obj.Kind == ast.Con {
//					vs, ok := obj.Decl.(*ast.ValueSpec)
//					if !ok {
//						beeLogger.Log.Fatalf("Unknown type without ValueSpec: %v\n", vs)
//					}
//
//					ti, ok := vs.Type.(*ast.Ident)
//					if !ok {
//						// TODO type inference, iota not support yet
//						continue
//					}
//					// Only add the enums that are defined by the current identifier
//					if ti.Name != k {
//						continue
//					}
//
//					// For all names and values, aggregate them by it's position so that we can sort them later.
//					for i, val := range vs.Values {
//						v, ok := val.(*ast.BasicLit)
//						if !ok {
//							beeLogger.Log.Warnf("Unknown type without BasicLit: %v\n", v)
//							continue
//						}
//						enums[int(val.Pos())] = fmt.Sprintf("%s = %s", vs.Names[i].Name, v.Value)
//						switch v.Kind {
//						case token.INT:
//							vv, err := strconv.Atoi(v.Value)
//							if err != nil {
//								beeLogger.Log.Warnf("Unknown type with BasicLit to int: %v\n", v.Value)
//								continue
//							}
//							enumValues[int(val.Pos())] = vv
//						case token.FLOAT:
//							vv, err := strconv.ParseFloat(v.Value, 64)
//							if err != nil {
//								beeLogger.Log.Warnf("Unknown type with BasicLit to int: %v\n", v.Value)
//								continue
//							}
//							enumValues[int(val.Pos())] = vv
//						default:
//							enumValues[int(val.Pos())] = strings.Trim(v.Value, `"`)
//						}
//
//					}
//				}
//			}
//		}
//	}
//	// Sort the enums by position
//	if len(enums) > 0 {
//		var keys []int
//		for k := range enums {
//			keys = append(keys, k)
//		}
//		sort.Ints(keys)
//		for _, k := range keys {
//			m.Enum = append(m.Enum, enums[k])
//		}
//		// Automatically use the first enum value as the example.
//		m.Example = enumValues[keys[0]]
//	}
//
//}
//
//func parseStruct(st *ast.StructType, k string, realTypes *[]string, astPkgs []*ast.Package, packageName string) {
//	m.Title = k
//	if st.Fields.List != nil {
//		m.Properties = make(map[string]swagger.Propertie)
//		for _, field := range st.Fields.List {
//			isSlice, realType, sType := typeAnalyser(field)
//			if (isSlice && isBasicType(realType)) || sType == "object" {
//				if len(strings.Split(realType, " ")) > 1 {
//					realType = strings.Replace(realType, " ", ".", -1)
//					realType = strings.Replace(realType, "&", "", -1)
//					realType = strings.Replace(realType, "{", "", -1)
//					realType = strings.Replace(realType, "}", "", -1)
//				} else {
//					realType = packageName + "." + realType
//				}
//			}
//			*realTypes = append(*realTypes, realType)
//			mp := swagger.Propertie{}
//			isObject := false
//			if isSlice {
//				mp.Type = "array"
//				if sType, ok := basicTypes[(strings.Replace(realType, "[]", "", -1))]; ok {
//					typeFormat := strings.Split(sType, ":")
//					mp.Items = &swagger.Propertie{
//						Type:   typeFormat[0],
//						Format: typeFormat[1],
//					}
//				} else {
//					mp.Items = &swagger.Propertie{
//						Ref: "#/definitions/" + realType,
//					}
//				}
//			} else {
//				if sType == "object" {
//					isObject = true
//					mp.Ref = "#/definitions/" + realType
//				} else if isBasicType(realType) {
//					typeFormat := strings.Split(sType, ":")
//					mp.Type = typeFormat[0]
//					mp.Format = typeFormat[1]
//				} else if realType == "map" {
//					typeFormat := strings.Split(sType, ":")
//					mp.AdditionalProperties = &swagger.Propertie{
//						Type:   typeFormat[0],
//						Format: typeFormat[1],
//					}
//				}
//			}
//			if field.Names != nil {
//
//				// set property name as field name
//				var name = field.Names[0].Name
//
//				// if no tag skip tag processing
//				if field.Tag == nil {
//					m.Properties[name] = mp
//					continue
//				}
//
//				var tagValues []string
//
//				stag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
//
//				defaultValue := stag.Get("doc")
//				if defaultValue != "" {
//					r, _ := regexp.Compile(`default\((.*)\)`)
//					if r.MatchString(defaultValue) {
//						res := r.FindStringSubmatch(defaultValue)
//						mp.Default = str2RealType(res[1], realType)
//
//					} else {
//						beeLogger.Log.Warnf("Invalid default value: %s", defaultValue)
//					}
//				}
//
//				tag := stag.Get("json")
//
//				if tag != "" {
//					tagValues = strings.Split(tag, ",")
//				}
//
//				// dont add property if json tag first value is "-"
//				if len(tagValues) == 0 || tagValues[0] != "-" {
//
//					// set property name to the left most json tag value only if is not omitempty
//					if len(tagValues) > 0 && tagValues[0] != "omitempty" {
//						name = tagValues[0]
//					}
//
//					if thrifttag := stag.Get("thrift"); thrifttag != "" {
//						ts := strings.Split(thrifttag, ",")
//						if ts[0] != "" {
//							name = ts[0]
//						}
//					}
//					if required := stag.Get("required"); required != "" {
//						m.Required = append(m.Required, name)
//					}
//					if desc := stag.Get("description"); desc != "" {
//						mp.Description = desc
//					}
//
//					if example := stag.Get("example"); example != "" && !isObject && !isSlice {
//						mp.Example = str2RealType(example, realType)
//					}
//
//					m.Properties[name] = mp
//				}
//				if ignore := stag.Get("ignore"); ignore != "" {
//					continue
//				}
//			} else {
//				// only parse case of when embedded field is TypeName
//				// cases of *TypeName and Interface are not handled, maybe useless for swagger spec
//				tag := ""
//				if field.Tag != nil {
//					stag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
//					tag = stag.Get("json")
//				}
//
//				if tag != "" {
//					tagValues := strings.Split(tag, ",")
//					if tagValues[0] == "-" {
//						//if json tag is "-", omit
//						continue
//					} else {
//						//if json tag is "something", output: something #definition/pkgname.Type
//						m.Properties[tagValues[0]] = mp
//						continue
//					}
//				} else {
//					//if no json tag, expand all fields of the type here
//					nm := &swagger.Schema{}
//					for _, pkg := range astPkgs {
//						for _, fl := range pkg.Files {
//							for nameOfObj, obj := range fl.Scope.Objects {
//								if obj.Name == fmt.Sprint(field.Type) {
//									parseObject(obj, nameOfObj, nm, realTypes, astPkgs, pkg.Name)
//								}
//							}
//						}
//					}
//					for name, p := range nm.Properties {
//						m.Properties[name] = p
//					}
//					continue
//				}
//			}
//		}
//	}
//}
//
//func typeAnalyser(f *ast.Field) (isSlice bool, realType, swaggerType string) {
//	if arr, ok := f.Type.(*ast.ArrayType); ok {
//		if isBasicType(fmt.Sprint(arr.Elt)) {
//			return true, fmt.Sprintf("[]%v", arr.Elt), basicTypes[fmt.Sprint(arr.Elt)]
//		}
//		if mp, ok := arr.Elt.(*ast.MapType); ok {
//			return false, fmt.Sprintf("map[%v][%v]", mp.Key, mp.Value), "object"
//		}
//		if star, ok := arr.Elt.(*ast.StarExpr); ok {
//			return true, fmt.Sprint(star.X), "object"
//		}
//		return true, fmt.Sprint(arr.Elt), "object"
//	}
//	switch t := f.Type.(type) {
//	case *ast.StarExpr:
//		basicType := fmt.Sprint(t.X)
//		if object, isStdLibObject := stdlibObject[basicType]; isStdLibObject {
//			basicType = object
//		}
//		if k, ok := basicTypes[basicType]; ok {
//			return false, basicType, k
//		}
//		return false, basicType, "object"
//	case *ast.MapType:
//		val := fmt.Sprintf("%v", t.Value)
//		if isBasicType(val) {
//			return false, "map", basicTypes[val]
//		}
//		return false, val, "object"
//	}
//	basicType := fmt.Sprint(f.Type)
//	if object, isStdLibObject := stdlibObject[basicType]; isStdLibObject {
//		basicType = object
//	}
//	if k, ok := basicTypes[basicType]; ok {
//		return false, basicType, k
//	}
//	return false, basicType, "object"
//}
