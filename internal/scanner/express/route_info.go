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
	Router        string // 이 라우트가 등록된 라우터 변수명 (prefix 해석용)
	Handler       string
	HandlerNode   *sitter.Node
	Middleware    []string
	Line          int
	ZodValidators []zod.ValidatorInfo
	AuthLevel     string
	Roles         []string
}
