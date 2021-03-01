package basicutils

import (
	"../logger"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func ListPublicFuncInPackage(packageName string) []*ast.FuncDecl {
	packageName = strings.ReplaceAll(packageName, ".", "/")

	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, packageName, nil, parser.ParseComments)
	if err != nil {
		logger.Error("Failed To Parse Package: %v", err)
		return nil
	}

	var funcList []*ast.FuncDecl
	for _, pack := range packs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					funcList = append(funcList, fn)
				}
			}
		}
	}
	return funcList
}
