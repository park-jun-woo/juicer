//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 토큰 위치가 속한 AST 파일을 패키지 목록에서 찾는다
package gogin

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

func findFileForPos(pos token.Pos, pkgs []*packages.Package) *ast.File {
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			if file.Pos() <= pos && pos <= file.End() {
				return file
			}
		}
	}
	return nil
}
