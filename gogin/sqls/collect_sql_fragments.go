//ff:func feature=sql type=parse control=sequence
//ff:what 메서드 body에서 SQL 키워드를 포함한 백틱 문자열 수집
package sqls

import (
	"go/ast"
	"go/token"
	"strings"
)

// collectSQLFragments collects backtick string literals from the method body
// that contain SQL keywords.
//
func collectSQLFragments(body *ast.BlockStmt) []string {
	if body == nil {
		return nil
	}

	var fragments []string
	ast.Inspect(body, func(n ast.Node) bool {
		lit, ok := n.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}
		val := lit.Value
		if !strings.HasPrefix(val, "`") {
			return true
		}
		// Strip backticks
		content := val[1 : len(val)-1]
		content = strings.TrimSpace(content)

		if len(content) < 10 {
			return true
		}
		if !sqlKeywords.MatchString(content) {
			return true
		}

		// Normalize whitespace for readability
		content = normalizeWhitespace(content)
		fragments = append(fragments, content)
		return true
	})
	return fragments
}

