//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 패키지 목록에서 Gin 라우트를 추출하고 정렬한다
package scanner

import (
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

func extractRoutes(pkgs []*packages.Package, root string) []Endpoint {
	var endpoints []Endpoint

	for _, pkg := range pkgs {
		for i, file := range pkg.Syntax {
			if i >= len(pkg.CompiledGoFiles) {
				continue
			}
			rel, _ := filepath.Rel(root, pkg.CompiledGoFiles[i])
			if strings.HasSuffix(rel, ".gen.go") {
				continue
			}
			eps := scanFile(file, rel, pkg.Fset)
			endpoints = append(endpoints, eps...)
		}
	}

	sort.Slice(endpoints, func(i, j int) bool {
		if endpoints[i].Path != endpoints[j].Path {
			return endpoints[i].Path < endpoints[j].Path
		}
		return endpoints[i].Method < endpoints[j].Method
	})
	return endpoints
}

