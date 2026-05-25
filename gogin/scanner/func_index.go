//ff:type feature=scan type=model
//ff:what funcIndex 데이터 구조
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
)

// funcIndex — 프로젝트 내 모든 함수/메서드 선언을 pos로 인덱싱
type funcIndex struct {
	byPos map[token.Pos]*ast.FuncDecl // 선언 위치 → FuncDecl
	info  map[token.Pos]*types.Info   // 선언 위치 → 해당 패키지의 TypesInfo
}
