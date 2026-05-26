//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 패키지 목록에서 Gin 라우트를 추출하고 정렬한다
package gogin

import (
	"go/ast"
	"path/filepath"
	"sort"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func extractRoutes(pkgs []*packages.Package, root string) ([]scanner.Endpoint, map[int][]ast.Expr) {
	var endpoints []scanner.Endpoint
	globalMap := map[int][]ast.Expr{}

	for _, pkg := range pkgs {
		for i, file := range pkg.Syntax {
			if i >= len(pkg.CompiledGoFiles) {
				continue
			}
			rel, _ := filepath.Rel(root, pkg.CompiledGoFiles[i])
			eps, hmap := scanFile(file, rel, pkg.Fset)
			// hmap 인덱스를 전역 오프셋으로 변환
			offset := len(endpoints)
			for k, v := range hmap {
				globalMap[offset+k] = v
			}
			endpoints = append(endpoints, eps...)
		}
	}

	// 정렬 전 인덱스 매핑을 유지하기 위해 정렬 전후 매핑 생성
	type indexed struct {
		ep  scanner.Endpoint
		idx int
	}
	items := make([]indexed, len(endpoints))
	for i, ep := range endpoints {
		items[i] = indexed{ep: ep, idx: i}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].ep.Path != items[j].ep.Path {
			return items[i].ep.Path < items[j].ep.Path
		}
		return items[i].ep.Method < items[j].ep.Method
	})

	sorted := make([]scanner.Endpoint, len(items))
	sortedMap := map[int][]ast.Expr{}
	for newIdx, item := range items {
		sorted[newIdx] = item.ep
		if exprs, ok := globalMap[item.idx]; ok {
			sortedMap[newIdx] = exprs
		}
	}

	return sorted, sortedMap
}
