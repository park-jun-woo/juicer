//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 상태 코드 인자 노드를 정수 또는 Response::HTTP_* 상수에서 코드 문자열로 해석한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func resolveStatusArg(arg *sitter.Node, src []byte) string {
	if code := constantStatusCode(arg, src); code != "" {
		return code
	}
	return strings.TrimSpace(nodeText(arg, src))
}
