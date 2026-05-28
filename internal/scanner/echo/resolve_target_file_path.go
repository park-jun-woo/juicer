//ff:func feature=scan type=extract control=sequence
//ff:what 토큰 위치에서 프로젝트 루트 기준 상대 파일 경로를 결정한다
package echo

import (
	"go/token"
	"path/filepath"
)

func resolveTargetFilePath(pos token.Pos, ctx *groupArgCtx) string {
	absPath := ctx.fset.Position(pos).Filename
	rel, err := filepath.Rel(ctx.root, absPath)
	if err != nil {
		return absPath
	}
	return rel
}
