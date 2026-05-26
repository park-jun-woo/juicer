//ff:func feature=sql type=parse control=sequence
//ff:what DB 호출의 인라인 문자열 인수에서 SQL 추출
package sqls

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

// collectInlineSQLArgs collects regular (double-quoted) string literals
// passed as arguments to ExecContext/QueryContext/QueryRowContext that contain SQL keywords.
//
func collectInlineSQLArgs(body *ast.BlockStmt) []string {
	if body == nil {
		return nil
	}

	var fragments []string
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		switch sel.Sel.Name {
		case "QueryContext", "QueryRowContext", "ExecContext":
		default:
			return true
		}

		for _, arg := range call.Args {
			lit, ok := arg.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				continue
			}
			val := lit.Value
			if strings.HasPrefix(val, "`") {
				continue // already handled by collectSQLFragments
			}
			// Double-quoted string — unquote using strconv for proper escape handling
			content, err := strconv.Unquote(val)
			if err != nil {
				content = val // fallback to original value
			}
			content = strings.TrimSpace(content)
			if len(content) < 10 {
				continue
			}
			if !sqlKeywords.MatchString(content) {
				continue
			}
			content = normalizeWhitespace(content)
			fragments = append(fragments, content)
		}
		return true
	})
	return fragments
}

