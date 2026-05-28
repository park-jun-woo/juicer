//ff:type feature=scan type=model topic=express
//ff:what 추출된 라우트 정보 구조체
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

type routeInfo struct {
	Method        string
	Path          string
	Handler       string
	HandlerNode   *sitter.Node
	Middleware    []string
	Line          int
	ZodValidators []zod.ValidatorInfo
	AuthLevel     string
	Roles         []string
}
