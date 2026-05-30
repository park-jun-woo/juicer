//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_FuncLit 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestAnalyzeExpr_FuncLit(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
var _ = func(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"a": "b"})
}
`
	lit := firstFuncLit(t, src)
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}

	analyzeExpr(ep, lit, nil, idx)
}
