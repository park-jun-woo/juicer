//ff:func feature=scan type=extract control=sequence
//ff:what 토큰 위치가 속한 파일에서 echo 패키지의 import alias를 결정한다
package echo

import (
	"go/token"
)

func resolveTargetEchoAlias(pos token.Pos, ctx *groupArgCtx) string {
	file := findFileForPos(pos, ctx.pkgs)
	if file == nil {
		return ""
	}
	return echoPkgName(file)
}
