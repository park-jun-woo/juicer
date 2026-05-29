//ff:func feature=scan type=extract control=sequence
//ff:what 문자열 연결 BinaryExpr에서 경로를 추출한다
package echo

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"
)

// extractBinaryPath extracts a path from a string concatenation BinaryExpr.
func extractBinaryPath(info *types.Info, e *ast.BinaryExpr) (string, bool) {
	if e.Op != token.ADD {
		return "", false
	}
	var parts []string
	collectStringParts(info, e, &parts)
	if len(parts) > 0 {
		return strings.Join(parts, ""), true
	}
	return "", false
}
