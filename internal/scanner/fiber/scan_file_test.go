//ff:func feature=scan type=test control=sequence
//ff:what scanFile — 파일 단위 라우트 추출 테스트
package fiber

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestScanFile_WithRoutes(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) {
	app.Get("/users", h)
	app.Post("/users", h)
}
var X = 1
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	eps, hmap := scanFile(file, "m.go", fset)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	_ = hmap
}

func TestScanFile_NoFiber(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", "package m\nfunc F() {}\n", 0)
	eps, hmap := scanFile(file, "m.go", fset)
	if eps != nil || hmap != nil {
		t.Fatalf("non-fiber file should yield nil, got %v %v", eps, hmap)
	}
}
