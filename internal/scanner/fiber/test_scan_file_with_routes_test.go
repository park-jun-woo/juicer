//ff:func feature=scan type=test control=sequence
//ff:what TestScanFile_WithRoutes 테스트
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
